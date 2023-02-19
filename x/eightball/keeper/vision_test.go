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

func createNVision(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Vision {
	items := make([]types.Vision, n)
	for i := range items {
		items[i].Id = keeper.AppendVision(ctx, items[i])
	}
	return items
}

func TestVisionGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVision(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVision(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVisionRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVision(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVision(ctx, item.Id)
		_, found := keeper.GetVision(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVisionGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVision(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVision(ctx)),
	)
}

func TestVisionCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVision(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVisionCount(ctx))
}
