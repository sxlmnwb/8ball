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

func createNSignatureShare(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SignatureShare {
	items := make([]types.SignatureShare, n)
	for i := range items {
		items[i].Id = keeper.AppendSignatureShare(ctx, items[i])
	}
	return items
}

func TestSignatureShareGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureShare(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSignatureShare(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSignatureShareRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureShare(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSignatureShare(ctx, item.Id)
		_, found := keeper.GetSignatureShare(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSignatureShareGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureShare(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSignatureShare(ctx)),
	)
}

func TestSignatureShareCount(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNSignatureShare(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSignatureShareCount(ctx))
}
