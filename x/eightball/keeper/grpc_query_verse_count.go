package keeper

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerseCount(goCtx context.Context, req *types.QueryVerseCountRequest) (*types.QueryVerseCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	magicKeyId, err := strconv.ParseUint(req.MagicKeyId, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	count := k.GetVerseCount(ctx, magicKeyId)
	countStr := strconv.FormatUint(count, 10)

	return &types.QueryVerseCountResponse{Count: countStr}, nil
}
