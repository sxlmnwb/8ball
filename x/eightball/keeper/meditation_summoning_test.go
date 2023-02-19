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

func createNMeditationSummoning(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MeditationSummoning {
	items := make([]types.MeditationSummoning, n)
	for i := range items {
		items[i].Id = keeper.AppendMeditationSummoning(ctx, items[i])
	}
	return items
}

func TestMeditationSummoningGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditationSummoning(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMeditationSummoning(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMeditationSummoningRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditationSummoning(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMeditationSummoning(ctx, item.Id)
		_, found := keeper.GetMeditationSummoning(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMeditationSummoningGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditationSummoning(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMeditationSummoning(ctx)),
	)
}

func TestMeditationSummoningCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditationSummoning(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMeditationSummoningCount(ctx))
}
