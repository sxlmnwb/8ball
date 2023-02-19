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

func createNImploring(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Imploring {
	items := make([]types.Imploring, n)
	for i := range items {
		items[i].Id = keeper.AppendImploring(ctx, items[i])
	}
	return items
}

func TestImploringGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNImploring(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetImploring(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestImploringRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNImploring(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveImploring(ctx, item.Id)
		_, found := keeper.GetImploring(ctx, item.Id)
		require.False(t, found)
	}
}

func TestImploringGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNImploring(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllImploring(ctx)),
	)
}

func TestImploringCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNImploring(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetImploringCount(ctx))
}
