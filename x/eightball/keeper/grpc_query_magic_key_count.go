package keeper

import (
	"context"

	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MagicKeyCount(goCtx context.Context, req *types.QueryMagicKeyCountRequest) (*types.QueryMagicKeyCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetMagicKeyCount(ctx)
	return &types.QueryMagicKeyCountResponse{Count: count}, nil
}
