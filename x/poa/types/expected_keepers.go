package types

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	IterateAllBalances(ctx context.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool))
}

// StakingKeeper defines the expected interface needed to retrieve account balances.
type StakingKeeper interface {
	GetParams(ctx context.Context) (stakingtypes.Params, error)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []stakingtypes.Validator)
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
	GetAllDelegations(ctx sdk.Context) (delegations []stakingtypes.Delegation)
	GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []stakingtypes.Delegation
	GetUnbondingDelegationsFromValidator(ctx sdk.Context, validator sdk.ValAddress) []stakingtypes.UnbondingDelegation
	SlashUnbondingDelegation(ctx sdk.Context, ubd stakingtypes.UnbondingDelegation, infractionHeight int64, slashFactor math.LegacyDec) (totalSlashAmount math.Int)
	RemoveDelegation(ctx sdk.Context, delegation stakingtypes.Delegation) error
	RemoveValidatorTokensAndShares(ctx sdk.Context, validator stakingtypes.Validator, sharesToRemove math.LegacyDec) (stakingtypes.Validator, math.Int)
	RemoveValidatorTokens(ctx sdk.Context, validator stakingtypes.Validator, tokensToRemove math.Int) stakingtypes.Validator
	BondDenom(ctx sdk.Context) string
}

type SlashingKeeper interface {
	GetParams(ctx sdk.Context) (params slashingtypes.Params)
}

type GovKeeper interface {
	SubmitProposal(ctx context.Context, messages []sdk.Msg, metadata string) (v1.Proposal, error)
}
