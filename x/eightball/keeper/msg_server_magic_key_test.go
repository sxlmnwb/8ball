package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"eightball/x/eightball/types"
)

func TestMagicKeyMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMagicKey(ctx, &types.MsgCreateMagicKey{})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}
