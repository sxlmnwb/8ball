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

func createNVerse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Verse {
	items := make([]types.Verse, n)
	for i := range items {
		items[i].Id = keeper.AppendVerse(ctx, items[i])
	}
	return items
}

func TestVerseGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVerse(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVerse(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVerseRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVerse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVerse(ctx, item.Id)
		_, found := keeper.GetVerse(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVerseGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVerse(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVerse(ctx)),
	)
}

func TestVerseCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNVerse(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVerseCount(ctx))
}
