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

func createNScriptureSignatureShare(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ScriptureSignatureShare {
	items := make([]types.ScriptureSignatureShare, n)
	for i := range items {
		items[i].Id = keeper.AppendScriptureSignatureShare(ctx, items[i])
	}
	return items
}

func TestScriptureSignatureShareGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNScriptureSignatureShare(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetScriptureSignatureShare(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestScriptureSignatureShareRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNScriptureSignatureShare(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScriptureSignatureShare(ctx, item.Id)
		_, found := keeper.GetScriptureSignatureShare(ctx, item.Id)
		require.False(t, found)
	}
}

func TestScriptureSignatureShareGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNScriptureSignatureShare(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllScriptureSignatureShare(ctx)),
	)
}

func TestScriptureSignatureShareCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNScriptureSignatureShare(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetScriptureSignatureShareCount(ctx))
}
