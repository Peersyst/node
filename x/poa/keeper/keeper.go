package keeper

import (
	"cosmossdk.io/math"
	"fmt"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	stakingkeeper "github.com/evmos/evmos/v20/x/staking/keeper"

	"github.com/xrplevm/node/v3/x/poa/types"
)

type (
	Keeper struct {
		cdc        codec.Codec
		paramstore paramtypes.Subspace
		authority  string                    // the address capable of executing a poa change. Usually the gov module account
		router     *baseapp.MsgServiceRouter // Msg server router
		bk         types.BankKeeper
		sk         stakingkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.Codec,
	ps paramtypes.Subspace,
	router *baseapp.MsgServiceRouter,
	bk types.BankKeeper,
	sk stakingkeeper.Keeper,
	authority string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	_, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		panic(err)
	}

	return &Keeper{
		cdc:        cdc,
		paramstore: ps,
		authority:  authority,
		router:     router,
		bk:         bk,
		sk:         sk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Router returns the gov keeper's router
func (k Keeper) Router() *baseapp.MsgServiceRouter {
	return k.router
}

func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) ExecuteAddValidator(ctx sdk.Context, msg *types.MsgAddValidator) error {
	// Check if the new validator already has staking power in the bank account
	accAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	valAddress := sdk.ValAddress(accAddress)
	if err != nil {
		return err
	}
	params, err := k.sk.GetParams(ctx)
	if err != nil {
		return err
	}
	denom := params.BondDenom
	balance := k.bk.GetBalance(ctx, accAddress, denom)
	if !balance.IsZero() {
		// Validator already has staking tokens in bank
		return types.ErrAddressHasBankTokens
	}

	// Check if the validator has bonded tokens in the staking module
	validator, err := k.sk.GetValidator(ctx, valAddress)
	if err == nil && !validator.Tokens.IsZero() {
		// Validator already has staking tokens bonded
		return types.ErrAddressHasBondedTokens
	}

	delegations, err := k.sk.GetAllDelegatorDelegations(ctx, accAddress)
	if err == nil {
		// Check if the delegations are greater than 0
		// Validator has delegations to other validators, not eligible for new tokens
		for _, delegation := range delegations {
			if !delegation.Shares.IsZero() {
				delVal, err := k.sk.GetValidator(ctx, sdk.ValAddress(delegation.GetValidatorAddr()))
				if err != nil {
					continue
				}
				if !delVal.Tokens.IsZero() {
					return types.ErrAddressHasDelegatedTokens
				}
			}
		}
	}

	// Check if address has unbonding delegations with balance
	// If so, return error since the account already has staking power
	unbondingBalance := math.ZeroInt()
	ubds, err := k.sk.GetUnbondingDelegationsFromValidator(ctx, valAddress)
	if err == nil {
		for _, ubd := range ubds {
			for _, entry := range ubd.Entries {
				unbondingBalance = unbondingBalance.Add(entry.Balance)
			}
		}
	}
	if !unbondingBalance.IsZero() {
		return types.ErrAddressHasUnbondingTokens
	}

	// All checks passed, mint new validator tokens and send them to the address
	coin := sdk.NewCoin(denom, sdk.DefaultPowerReduction)
	coins := sdk.NewCoins(coin)
	err = k.bk.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return err
	}
	err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coins)
	if err != nil {
		return err
	}

	pubKey, ok := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return errors.Wrapf(sdkerrors.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pubKey)
	}
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(
		valAddress.String(),
		pubKey,
		coin,
		msg.Description,
		stakingtypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
		math.OneInt(),
	)
	if err != nil {
		return err
	}
	handler := k.Router().Handler(createValidatorMsg)
	_, err = handler(ctx, createValidatorMsg)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddValidator,
			sdk.NewAttribute(types.AttributeValidator, accAddress.String()),
			sdk.NewAttribute(types.AttributeHeight, fmt.Sprintf("%d", ctx.BlockHeight())),
			sdk.NewAttribute(types.AttributeStakingTokens, fmt.Sprintf("%d", validator.Tokens)),
			sdk.NewAttribute(types.AttributeBankTokens, balance.String()),
		),
	)

	return nil
}

func (k Keeper) ExecuteRemoveValidator(ctx sdk.Context, validatorAddress string) error {
	accAddress, err := sdk.AccAddressFromBech32(validatorAddress)
	if err != nil {
		return err
	}
	params, err := k.sk.GetParams(ctx)
	if err != nil {
		return err
	}
	denom := params.BondDenom
	valAddress := sdk.ValAddress(accAddress)

	balance := k.bk.GetBalance(ctx, accAddress, denom)
	if balance.IsZero() {
		// Address has no balance in bank and is not a validator either
		// NOTE: Since delegations are not enabled in this version, we don't need to check for them
		return types.ErrAddressHasNoTokens
	} else {
		// If address also has tokens in the bank, we need to remove them and burn them
		coins := sdk.NewCoins(balance)
		err = k.bk.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, coins)
		if err != nil {
			return err
		}

		err = k.bk.BurnCoins(ctx, types.ModuleName, coins)
		if err != nil {
			return err
		}
	}

	// Check if address is a validator or has balance some balance in bank
	validator, err := k.sk.GetValidator(ctx, valAddress)
	if err != nil {
		ctx.Logger().Warn("Error fetching validator", "error", err)
		// Not failing hard here, since the address could be a delegator without a validator
		return nil
	}

	// If address is a validator, we need to check if there are unbonding delegations currently
	// and slash them. We also need to remove all the tokens from the validator and burn them
	// from the staking module account
	ubds, err := k.sk.GetUnbondingDelegationsFromValidator(ctx, valAddress)
	if err != nil {
		ctx.Logger().Error("Error fetching unbonding delegations", "error", err)
	} else {
		for _, ubd := range ubds {
			totalSlashAmount, err := k.sk.SlashUnbondingDelegation(ctx, ubd, 0, math.LegacyOneDec())
			if err != nil {
				// Fail hard due to error when to slashing a validator's unbonding delegations
				return err
			}
			ctx.Logger().Info("Slashed unbonding delegation", "validator", valAddress, "amount", totalSlashAmount)
		}
	}

	// Remove delegator shares
	delegations, err := k.sk.GetAllDelegatorDelegations(ctx, accAddress)
	if err != nil {
		ctx.Logger().Error("Error fetching delegations", "error", err)
	} else {
		for _, delegation := range delegations {
			if !delegation.Shares.IsZero() {
				if err := k.sk.RemoveDelegation(ctx, delegation); err != nil {
					return err
				}
			}
		}
	}

	validator, removedTokens, err := k.sk.RemoveValidatorTokensAndShares(ctx, validator, validator.DelegatorShares)
	if err != nil {
		// Fail hard due to error when removing validator tokens and shares
		return err
	}
	ctx.Logger().Info("Removed validator tokens", "validator", validatorAddress, "tokens", removedTokens)
	changedVal, err := k.sk.RemoveValidatorTokens(ctx, validator, validator.Tokens)
	if err != nil {
		// Fail hard due to error when removing validator tokens
		return err
	}

	switch changedVal.GetStatus() {
	case stakingtypes.Bonded:
		coins := sdk.NewCoins(sdk.NewCoin(denom, validator.Tokens))
		err = k.bk.BurnCoins(ctx, stakingtypes.BondedPoolName, coins)
		if err != nil {
			return err
		}
	case stakingtypes.Unbonding, stakingtypes.Unbonded:
		coins := sdk.NewCoins(sdk.NewCoin(denom, validator.Tokens))
		err = k.bk.BurnCoins(ctx, stakingtypes.NotBondedPoolName, coins)
		if err != nil {
			return err
		}
	default:
		return types.ErrInvalidValidatorStatus
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveValidator,
			sdk.NewAttribute(types.AttributeValidator, validatorAddress),
			sdk.NewAttribute(types.AttributeHeight, fmt.Sprintf("%d", ctx.BlockHeight())),
			sdk.NewAttribute(types.AttributeStakingTokens, fmt.Sprintf("%d", validator.Tokens)),
			sdk.NewAttribute(types.AttributeBankTokens, balance.String()),
		),
	)

	return nil
}
