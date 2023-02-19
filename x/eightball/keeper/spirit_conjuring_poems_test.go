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

func createNSpiritConjuringPoems(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SpiritConjuringPoems {
	items := make([]types.SpiritConjuringPoems, n)
	for i := range items {
		items[i].Id = keeper.AppendSpiritConjuringPoems(ctx, items[i])
	}
	return items
}

func TestSpiritConjuringPoemsGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSpiritConjuringPoems(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSpiritConjuringPoems(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSpiritConjuringPoemsRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSpiritConjuringPoems(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSpiritConjuringPoems(ctx, item.Id)
		_, found := keeper.GetSpiritConjuringPoems(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSpiritConjuringPoemsGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSpiritConjuringPoems(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSpiritConjuringPoems(ctx)),
	)
}

func TestSpiritConjuringPoemsCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSpiritConjuringPoems(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSpiritConjuringPoemsCount(ctx))
}
