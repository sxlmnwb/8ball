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

func createNEncryptedPreSign(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EncryptedPreSign {
	items := make([]types.EncryptedPreSign, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEncryptedPreSign(ctx, items[i])
	}
	return items
}

func TestEncryptedPreSignGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedPreSign(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEncryptedPreSign(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEncryptedPreSignRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedPreSign(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEncryptedPreSign(ctx,
			item.Index,
		)
		_, found := keeper.GetEncryptedPreSign(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEncryptedPreSignGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedPreSign(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEncryptedPreSign(ctx)),
	)
}
