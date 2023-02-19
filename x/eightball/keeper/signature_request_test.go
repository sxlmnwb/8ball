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

func createNSignatureRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SignatureRequest {
	items := make([]types.SignatureRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendSignatureRequest(ctx, items[i])
	}
	return items
}

func TestSignatureRequestGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSignatureRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSignatureRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSignatureRequest(ctx, item.Id)
		_, found := keeper.GetSignatureRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSignatureRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSignatureRequest(ctx)),
	)
}

func TestSignatureRequestCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSignatureRequestCount(ctx))
}
