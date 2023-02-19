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

func createNSignedMessage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SignedMessage {
	items := make([]types.SignedMessage, n)
	for i := range items {
		items[i].Id = keeper.AppendSignedMessage(ctx, items[i])
	}
	return items
}

func TestSignedMessageGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedMessage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSignedMessage(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSignedMessageRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedMessage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSignedMessage(ctx, item.Id)
		_, found := keeper.GetSignedMessage(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSignedMessageGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedMessage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSignedMessage(ctx)),
	)
}

func TestSignedMessageCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignedMessage(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSignedMessageCount(ctx))
}
