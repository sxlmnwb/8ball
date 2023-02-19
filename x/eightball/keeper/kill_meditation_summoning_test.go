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

func createNKillMeditationSummoning(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KillMeditationSummoning {
	items := make([]types.KillMeditationSummoning, n)
	for i := range items {
		items[i].Id = keeper.AppendKillMeditationSummoning(ctx, items[i])
	}
	return items
}

func TestKillMeditationSummoningGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMeditationSummoning(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKillMeditationSummoning(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKillMeditationSummoningRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMeditationSummoning(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKillMeditationSummoning(ctx, item.Id)
		_, found := keeper.GetKillMeditationSummoning(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKillMeditationSummoningGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMeditationSummoning(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKillMeditationSummoning(ctx)),
	)
}

func TestKillMeditationSummoningCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMeditationSummoning(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKillMeditationSummoningCount(ctx))
}
