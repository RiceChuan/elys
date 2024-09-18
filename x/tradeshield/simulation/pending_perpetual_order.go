package simulation

import (
	"math/rand"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/elys-network/elys/x/tradeshield/keeper"
	"github.com/elys-network/elys/x/tradeshield/types"
)

func SimulateMsgCreatePendingPerpetualOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgCreatePendingPerpetualOrder{
			Creator: simAccount.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdatePendingPerpetualOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount               = simtypes.Account{}
			pendingPerpetualOrder    = types.PerpetualOrder{}
			msg                      = &types.MsgUpdatePendingPerpetualOrder{}
			allPendingPerpetualOrder = k.GetAllPendingPerpetualOrder(ctx)
			found                    = false
		)
		for _, obj := range allPendingPerpetualOrder {
			simAccount, found = FindAccount(accs, obj.OwnerAddress)
			if found {
				pendingPerpetualOrder = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "pendingPerpetualOrder creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Id = pendingPerpetualOrder.OrderId

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeletePendingPerpetualOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount               = simtypes.Account{}
			pendingPerpetualOrder    = types.PerpetualOrder{}
			msg                      = &types.MsgUpdatePendingPerpetualOrder{}
			allPendingPerpetualOrder = k.GetAllPendingPerpetualOrder(ctx)
			found                    = false
		)
		for _, obj := range allPendingPerpetualOrder {
			simAccount, found = FindAccount(accs, obj.OwnerAddress)
			if found {
				pendingPerpetualOrder = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "pendingPerpetualOrder creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Id = pendingPerpetualOrder.OrderId

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}