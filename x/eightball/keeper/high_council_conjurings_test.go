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

func createNHighCouncilConjurings(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HighCouncilConjurings {
	items := make([]types.HighCouncilConjurings, n)
	for i := range items {
		items[i].Id = keeper.AppendHighCouncilConjurings(ctx, items[i])
	}
	return items
}

func TestHighCouncilConjuringsGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncilConjurings(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHighCouncilConjurings(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHighCouncilConjuringsRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncilConjurings(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHighCouncilConjurings(ctx, item.Id)
		_, found := keeper.GetHighCouncilConjurings(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHighCouncilConjuringsGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncilConjurings(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHighCouncilConjurings(ctx)),
	)
}

func TestHighCouncilConjuringsCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNHighCouncilConjurings(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHighCouncilConjuringsCount(ctx))
}
