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

func createNKillMagicKeySummoning(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KillMagicKeySummoning {
	items := make([]types.KillMagicKeySummoning, n)
	for i := range items {
		items[i].Id = keeper.AppendKillMagicKeySummoning(ctx, items[i])
	}
	return items
}

func TestKillMagicKeySummoningGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMagicKeySummoning(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKillMagicKeySummoning(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKillMagicKeySummoningRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMagicKeySummoning(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKillMagicKeySummoning(ctx, item.Id)
		_, found := keeper.GetKillMagicKeySummoning(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKillMagicKeySummoningGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMagicKeySummoning(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKillMagicKeySummoning(ctx)),
	)
}

func TestKillMagicKeySummoningCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillMagicKeySummoning(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKillMagicKeySummoningCount(ctx))
}
