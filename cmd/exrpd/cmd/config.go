package cmd

import (
	"github.com/Peersyst/exrp/v2/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermint "github.com/evmos/evmos/v15/types"
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.SetCoinType(app.Bip44CoinType)
	config.SetPurpose(sdk.Purpose) // Shared
	// config.SetFullFundraiserPath(ethermint.BIP44HDPath) // nolint: staticcheck
	config.Seal()
}

const (
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "xrp"
	// BaseDenom defines to the default denomination used in EVM
	BaseDenom = "axrp"
)

func registerDenoms() {
	if err := sdk.RegisterDenom(DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(BaseDenom, sdk.NewDecWithPrec(1, ethermint.BaseDenomUnit)); err != nil {
		panic(err)
	}
}
