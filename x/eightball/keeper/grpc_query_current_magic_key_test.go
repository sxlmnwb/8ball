package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball/types"
)

func TestCurrentMagicKeyQuery(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestCurrentMagicKey(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCurrentMagicKeyRequest
		response *types.QueryGetCurrentMagicKeyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCurrentMagicKeyRequest{},
			response: &types.QueryGetCurrentMagicKeyResponse{CurrentMagicKey: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CurrentMagicKey(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
