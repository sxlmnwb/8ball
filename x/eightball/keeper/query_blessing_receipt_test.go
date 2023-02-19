package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBlessingReceiptQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBlessingReceipt(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBlessingReceiptRequest
		response *types.QueryGetBlessingReceiptResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBlessingReceiptRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetBlessingReceiptResponse{BlessingReceipt: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBlessingReceiptRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetBlessingReceiptResponse{BlessingReceipt: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBlessingReceiptRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BlessingReceipt(wctx, tc.request)
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

func TestBlessingReceiptQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBlessingReceipt(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBlessingReceiptRequest {
		return &types.QueryAllBlessingReceiptRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BlessingReceiptAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BlessingReceipt), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BlessingReceipt),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BlessingReceiptAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BlessingReceipt), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BlessingReceipt),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BlessingReceiptAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.BlessingReceipt),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BlessingReceiptAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
