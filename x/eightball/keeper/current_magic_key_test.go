package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
)

func createTestCurrentMagicKey(keeper *keeper.Keeper, ctx sdk.Context) types.CurrentMagicKey {
	item := types.CurrentMagicKey{}
	keeper.SetCurrentMagicKey(ctx, item)
	return item
}

func TestCurrentMagicKeyGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	item := createTestCurrentMagicKey(keeper, ctx)
	rst, found := keeper.GetCurrentMagicKey(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestCurrentMagicKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	createTestCurrentMagicKey(keeper, ctx)
	keeper.RemoveCurrentMagicKey(ctx)
	_, found := keeper.GetCurrentMagicKey(ctx)
	require.False(t, found)
}
