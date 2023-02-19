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

func createNBlessing(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Blessing {
	items := make([]types.Blessing, n)
	for i := range items {
		items[i].Id = keeper.AppendBlessing(ctx, items[i])
	}
	return items
}

func TestBlessingGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessing(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBlessing(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBlessingRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessing(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlessing(ctx, item.Id)
		_, found := keeper.GetBlessing(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBlessingGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessing(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBlessing(ctx)),
	)
}

func TestBlessingCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessing(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBlessingCount(ctx))
}
