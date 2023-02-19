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

func createNMagicKey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MagicKey {
	items := make([]types.MagicKey, n)
	for i := range items {
		items[i].Id = keeper.AppendMagicKey(ctx, items[i])
	}
	return items
}

func TestMagicKeyGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKey(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMagicKey(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMagicKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKey(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMagicKey(ctx, item.Id)
		_, found := keeper.GetMagicKey(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMagicKeyGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKey(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMagicKey(ctx)),
	)
}

func TestMagicKeyCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNMagicKey(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMagicKeyCount(ctx))
}
