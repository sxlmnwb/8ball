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

func createNMagicKeySummoning(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MagicKeySummoning {
	items := make([]types.MagicKeySummoning, n)
	for i := range items {
		items[i].Id = keeper.AppendMagicKeySummoning(ctx, items[i])
	}
	return items
}

func TestMagicKeySummoningGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKeySummoning(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMagicKeySummoning(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMagicKeySummoningRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKeySummoning(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMagicKeySummoning(ctx, item.Id)
		_, found := keeper.GetMagicKeySummoning(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMagicKeySummoningGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKeySummoning(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMagicKeySummoning(ctx)),
	)
}

func TestMagicKeySummoningCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKeySummoning(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMagicKeySummoningCount(ctx))
}
