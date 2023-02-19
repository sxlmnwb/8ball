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

func createNSignedScripture(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SignedScripture {
	items := make([]types.SignedScripture, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSignedScripture(ctx, items[i])
	}
	return items
}

func TestSignedScriptureGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedScripture(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSignedScripture(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSignedScriptureRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedScripture(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSignedScripture(ctx,
			item.Index,
		)
		_, found := keeper.GetSignedScripture(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSignedScriptureGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedScripture(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSignedScripture(ctx)),
	)
}
