package keeper

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VisionAll(c context.Context, req *types.QueryAllVisionRequest) (*types.QueryAllVisionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	magKeyIdStr := strconv.FormatUint(req.MagicKeyId, 10)

	var visions []types.Vision
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	visionStore := prefix.NewStore(store, types.KeyPrefix(types.VisionKey+magKeyIdStr))

	pageRes, err := query.Paginate(visionStore, req.Pagination, func(key []byte, value []byte) error {
		var vision types.Vision
		if err := k.cdc.Unmarshal(value, &vision); err != nil {
			return err
		}

		visions = append(visions, vision)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVisionResponse{Vision: visions, Pagination: pageRes}, nil
}

func (k Keeper) Vision(c context.Context, req *types.QueryGetVisionRequest) (*types.QueryGetVisionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vision, found := k.GetVision(ctx, req.MagicKeyId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVisionResponse{Vision: vision}, nil
}
