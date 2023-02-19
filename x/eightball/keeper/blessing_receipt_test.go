package keeper_test

import (
	"strconv"
	"testing"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBlessingReceipt(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BlessingReceipt {
	items := make([]types.BlessingReceipt, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBlessingReceipt(ctx, items[i])
	}
	return items
}

func TestBlessingReceiptGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessingReceipt(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBlessingReceipt(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBlessingReceiptRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessingReceipt(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlessingReceipt(ctx,
			item.Index,
		)
		_, found := keeper.GetBlessingReceipt(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBlessingReceiptGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNBlessingReceipt(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBlessingReceipt(ctx)),
	)
}
