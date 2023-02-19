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

func createNHighCouncil(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HighCouncil {
	items := make([]types.HighCouncil, n)
	for i := range items {
		items[i].Id = keeper.AppendHighCouncil(ctx, items[i])
	}
	return items
}

func TestHighCouncilGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncil(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHighCouncil(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHighCouncilRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncil(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHighCouncil(ctx, item.Id)
		_, found := keeper.GetHighCouncil(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHighCouncilGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncil(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHighCouncil(ctx)),
	)
}

func TestHighCouncilCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncil(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHighCouncilCount(ctx))
}
