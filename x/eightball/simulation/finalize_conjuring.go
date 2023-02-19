package simulation

import (
	"math/rand"

	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFinalizeConjuring(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFinalizeConjuring{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the FinalizeConjuring simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "FinalizeConjuring simulation not implemented"), nil, nil
	}
}
