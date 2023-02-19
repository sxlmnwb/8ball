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

func createNEncryptedMagicKeyShare(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EncryptedMagicKeyShare {
	items := make([]types.EncryptedMagicKeyShare, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEncryptedMagicKeyShare(ctx, items[i])
	}
	return items
}

func TestEncryptedMagicKeyShareGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedMagicKeyShare(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEncryptedMagicKeyShare(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEncryptedMagicKeyShareRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedMagicKeyShare(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEncryptedMagicKeyShare(ctx,
			item.Index,
		)
		_, found := keeper.GetEncryptedMagicKeyShare(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEncryptedMagicKeyShareGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNEncryptedMagicKeyShare(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEncryptedMagicKeyShare(ctx)),
	)
}
