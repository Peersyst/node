package app_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ethante "github.com/evmos/evmos/v19/app/ante/evm"
	"github.com/evmos/evmos/v19/crypto/ethsecp256k1"
	poaante "github.com/xrplevm/node/v3/x/poa/ante"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/evmos/evmos/v19/app/ante"
	evmostypes "github.com/evmos/evmos/v19/types"
	"github.com/stretchr/testify/require"
	"github.com/xrplevm/node/v3/app"
)

func init() {
	simcli.GetSimulatorFlags()
}

const SimAppChainID = "simulation_777-1"

// NewSimApp disable feemarket on native tx, otherwise the cosmos-sdk simulation tests will fail.
func NewSimApp(logger log.Logger, db dbm.DB, config simulationtypes.Config) (*app.App, error) {
	encodingConfig := app.MakeEncodingConfig()
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	bApp := app.New(
		logger,
		db,
		nil,
		false,
		map[int64]bool{},
		app.DefaultNodeHome,
		simcli.FlagPeriodValue, encodingConfig,
		appOptions,
		baseapp.SetChainID(config.ChainID),
	)
	// disable feemarket on native tx
	options := ante.HandlerOptions{
		Cdc:                    encodingConfig.Codec,
		AccountKeeper:          bApp.AccountKeeper,
		BankKeeper:             bApp.BankKeeper,
		ExtensionOptionChecker: evmostypes.HasDynamicFeeExtensionOption,
		EvmKeeper:              bApp.EvmKeeper,
		FeegrantKeeper:         bApp.FeeGrantKeeper,
		IBCKeeper:              bApp.IBCKeeper,
		FeeMarketKeeper:        bApp.FeeMarketKeeper,
		SignModeHandler:        encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:         ante.SigVerificationGasConsumer,
		MaxTxGasWanted:         0,
		TxFeeChecker:           ethante.NewDynamicFeeChecker(bApp.EvmKeeper),
		StakingKeeper:          bApp.StakingKeeper,
		DistributionKeeper:     bApp.DistrKeeper,
		ExtraDecorator:         poaante.NewPoaDecorator(),
		AuthzDisabledMsgTypes: []string{
			sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
			sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
		},
	}

	if err := options.Validate(); err != nil {
		panic(err)
	}

	bApp.SetAnteHandler(ante.NewAnteHandler(options))
	if err := bApp.LoadLatestVersion(); err != nil {
		return nil, err
	}
	return bApp, nil
}

// RandomAccounts generates n random accounts
func RandomAccounts(r *rand.Rand, n int) []simulationtypes.Account {
	accs := make([]simulationtypes.Account, n)

	for i := 0; i < n; i++ {
		// don't need that much entropy for simulation
		privkeySeed := make([]byte, 32)
		r.Read(privkeySeed)

		accs[i].PrivKey = &ethsecp256k1.PrivKey{Key: privkeySeed}
		accs[i].PubKey = accs[i].PrivKey.PubKey()
		accs[i].Address = sdk.AccAddress(accs[i].PubKey.Address())

		accs[i].ConsKey = ed25519.GenPrivKeyFromSecret(privkeySeed)
	}

	return accs
}

// BenchmarkSimulation run the chain simulation
// Running using starport command:
// `ignite chain simulate -v --numBlocks 200 --blockSize 50`
// Running as go benchmark test:
// `go test -benchmem -run=^$ -bench ^BenchmarkSimulation ./app -NumBlocks=200 -BlockSize 50 -Commit=true -Verbose=true -Enabled=true`
//
//nolint:dupl
func BenchmarkSimulation(b *testing.B) {
	simcli.FlagSeedValue = time.Now().Unix()
	simcli.FlagVerboseValue = true
	simcli.FlagCommitValue = true
	simcli.FlagEnabledValue = true

	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID
	db, dir, logger, _, err := simtestutil.SetupSimulation(
		config,
		"leveldb-bApp-sim",
		"Simulation",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)

	require.NoError(b, err, "simulation setup failed")

	config.ChainID = SimAppChainID

	b.Cleanup(func() {
		require.NoError(b, db.Close())
		require.NoError(b, os.RemoveAll(dir))
	})

	bApp, _ := NewSimApp(logger, db, config)

	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		bApp.BaseApp,
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		RandomAccounts,
		simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
		bApp.ModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(bApp, config, simParams)
	require.NoError(b, err)
	require.NoError(b, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}
}

//nolint:dupl
func TestFullAppSimulation(t *testing.T) {
	simcli.FlagSeedValue = time.Now().Unix()
	simcli.FlagVerboseValue = true
	simcli.FlagCommitValue = true
	simcli.FlagEnabledValue = true

	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID
	db, dir, logger, _, err := simtestutil.SetupSimulation(
		config,
		"leveldb-bApp-sim",
		"Simulation",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)

	require.NoError(t, err, "simulation setup failed")

	config.ChainID = SimAppChainID

	t.Cleanup(func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	})

	bApp, _ := NewSimApp(logger, db, config)

	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		bApp.BaseApp,
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		RandomAccounts,
		simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
		bApp.ModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(bApp, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}
}
