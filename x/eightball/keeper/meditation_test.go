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

func createNMeditation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Meditation {
	items := make([]types.Meditation, n)
	for i := range items {
		items[i].Id = keeper.AppendMeditation(ctx, items[i])
	}
	return items
}

func TestMeditationGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditation(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMeditation(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMeditationRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMeditation(ctx, item.Id)
		_, found := keeper.GetMeditation(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMeditationGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMeditation(ctx)),
	)
}

func TestMeditationCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMeditation(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMeditationCount(ctx))
}
