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

func createNKillConjuring(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KillConjuring {
	items := make([]types.KillConjuring, n)
	for i := range items {
		items[i].Id = keeper.AppendKillConjuring(ctx, items[i])
	}
	return items
}

func TestKillConjuringGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillConjuring(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKillConjuring(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKillConjuringRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillConjuring(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKillConjuring(ctx, item.Id)
		_, found := keeper.GetKillConjuring(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKillConjuringGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillConjuring(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKillConjuring(ctx)),
	)
}

func TestKillConjuringCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillConjuring(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKillConjuringCount(ctx))
}
