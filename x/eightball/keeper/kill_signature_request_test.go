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

func createNKillSignatureRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KillSignatureRequest {
	items := make([]types.KillSignatureRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendKillSignatureRequest(ctx, items[i])
	}
	return items
}

func TestKillSignatureRequestGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKillSignatureRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKillSignatureRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillSignatureRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKillSignatureRequest(ctx, item.Id)
		_, found := keeper.GetKillSignatureRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKillSignatureRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillSignatureRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKillSignatureRequest(ctx)),
	)
}

func TestKillSignatureRequestCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNKillSignatureRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKillSignatureRequestCount(ctx))
}
