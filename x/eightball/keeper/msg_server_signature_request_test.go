package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"eightball/x/eightball/types"
)

func TestSignatureRequestMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateSignatureRequest(ctx, &types.MsgCreateSignatureRequest{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}
