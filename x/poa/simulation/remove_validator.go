package simulation

import (
	"math/rand"

	"github.com/Peersyst/exrp/v2/x/poa/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRemoveValidator() simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgRemoveValidator{}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RemoveValidator simulation not implemented"), nil, nil
	}
}
