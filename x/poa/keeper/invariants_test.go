package keeper

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/xrplevm/node/v3/x/poa/testutil"
)

func TestStakingPowerInvariant_Valid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {
			stakingKeeper.EXPECT().GetAllValidators(ctx).Return([]stakingtypes.Validator{
				{
					Tokens: sdk.DefaultPowerReduction,
				},
				{
					Tokens: sdk.ZeroInt(),
				},
			})
		},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {},
	)

	invariant := StakingPowerInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.False(t, broken, msg)
}

func TestStakingPowerInvariant_Invalid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {
			stakingKeeper.EXPECT().GetAllValidators(ctx).Return([]stakingtypes.Validator{
				{
					Tokens: sdk.DefaultPowerReduction,
				},
				{
					Tokens: sdk.DefaultPowerReduction.Add(sdk.OneInt()),
				},
			})
		},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {},
	)

	invariant := StakingPowerInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.True(t, broken, msg)
}

func TestSelfDelegationInvariant_Valid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {
			stakingKeeper.EXPECT().GetAllDelegations(ctx).Return([]stakingtypes.Delegation{
				{
					DelegatorAddress: "ethm13ued6aqj3w7jvks4l270dunhue0a9y7tspnpn5",
					ValidatorAddress: "ethmvaloper13ued6aqj3w7jvks4l270dunhue0a9y7tl3edtf",
				},
				{
					DelegatorAddress: "ethm13ued6aqj3w7jvks4l270dunhue0a9y7tspnpn5",
					ValidatorAddress: "ethmvaloper13ued6aqj3w7jvks4l270dunhue0a9y7tl3edtf",
				},
			})
		},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {},
	)

	invariant := SelfDelegationInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.False(t, broken, msg)
}

func TestSelfDelegationInvariant_Invalid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {
			stakingKeeper.EXPECT().GetAllDelegations(ctx).Return([]stakingtypes.Delegation{
				{
					DelegatorAddress: "ethm1wunfhl05vc8r8xxnnp8gt62wa54r6y52pg03zq",
					ValidatorAddress: "ethmvaloper13ued6aqj3w7jvks4l270dunhue0a9y7tl3edtf",
				},
			})
		},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {},
	)

	invariant := SelfDelegationInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.True(t, broken, msg)
}

func TestCheckSlashingParamsInvariant_Valid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {
			slashingKeeper.EXPECT().GetParams(ctx).Return(slashingtypes.Params{
				SlashFractionDoubleSign: math.LegacyZeroDec(),
				SlashFractionDowntime:   math.LegacyZeroDec(),
			})
		},
	)

	invariant := CheckSlashingParamsInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.False(t, broken, msg)
}

func TestCheckSlashingParamsInvariant_Invalid(t *testing.T) {
	poaKeeper, ctx := setupPoaKeeper(
		t,
		func(ctx sdk.Context, stakingKeeper *testutil.MockStakingKeeper) {},
		func(ctx sdk.Context, bankKeeper *testutil.MockBankKeeper) {},
		func(ctx sdk.Context, slashingKeeper *testutil.MockSlashingKeeper) {
			slashingKeeper.EXPECT().GetParams(ctx).Return(slashingtypes.Params{
				SignedBlocksWindow:      100,
				MinSignedPerWindow:      sdk.NewDecWithPrec(5, 1), // 0.5
				DowntimeJailDuration:    time.Duration(10 * time.Minute),
				SlashFractionDoubleSign: sdk.NewDecWithPrec(5, 2), // 0.05
				SlashFractionDowntime:   sdk.NewDecWithPrec(6, 1), // 0.6 (invalid, should be less than MinSignedPerWindow)
			})
		},
	)

	invariant := CheckSlashingParamsInvariant(*poaKeeper)
	msg, broken := invariant(ctx)
	require.True(t, broken, msg)
}
