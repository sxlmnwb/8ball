package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "eightball/testutil/keeper"
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestScriptureMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EightballKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	witness := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateScripture{Witness: witness,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateScripture(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetScripture(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Witness, rst.Address)
	}
}
