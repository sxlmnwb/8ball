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

func TestSignedScriptureMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EightballKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSignedScripture{Witness: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateSignedScripture(wctx, expected)
		require.NoError(t, err)
		_, found := k.GetSignedScripture(ctx,
			expected.Index,
		)
		require.True(t, found)
	}
}
