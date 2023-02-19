package keeper_test

import (
	"testing"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNKillScriptureSignatureRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KillScriptureSignatureRequest {
	items := make([]types.KillScriptureSignatureRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendKillScriptureSignatureRequest(ctx, items[i])
	}
	return items
}

func TestKillScriptureSignatureRequestGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillScriptureSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKillScriptureSignatureRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKillScriptureSignatureRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillScriptureSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKillScriptureSignatureRequest(ctx, item.Id)
		_, found := keeper.GetKillScriptureSignatureRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKillScriptureSignatureRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillScriptureSignatureRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKillScriptureSignatureRequest(ctx)),
	)
}

func TestKillScriptureSignatureRequestCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillScriptureSignatureRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKillScriptureSignatureRequestCount(ctx))
}
