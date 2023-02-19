package keeper

import (
	"context"

	"eightball/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CurrentMagicKey(c context.Context, req *types.QueryGetCurrentMagicKeyRequest) (*types.QueryGetCurrentMagicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCurrentMagicKey(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCurrentMagicKeyResponse{CurrentMagicKey: val}, nil
}
