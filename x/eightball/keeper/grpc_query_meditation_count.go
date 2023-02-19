package keeper

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MeditationCount(goCtx context.Context, req *types.QueryMeditationCountRequest) (*types.QueryMeditationCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetMeditationCount(ctx, req.MagicKeyId)
	countStr := strconv.FormatUint(count, 10)

	return &types.QueryMeditationCountResponse{Count: countStr}, nil
}
