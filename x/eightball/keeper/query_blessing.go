package keeper

import (
	"context"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlessingAll(c context.Context, req *types.QueryAllBlessingRequest) (*types.QueryAllBlessingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blessings []types.Blessing
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	blessingStore := prefix.NewStore(store, types.KeyPrefix(types.BlessingKey+req.Index))

	pageRes, err := query.Paginate(blessingStore, req.Pagination, func(key []byte, value []byte) error {
		var blessing types.Blessing
		if err := k.cdc.Unmarshal(value, &blessing); err != nil {
			return err
		}

		blessings = append(blessings, blessing)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlessingResponse{Blessing: blessings, Pagination: pageRes}, nil
}

func (k Keeper) Blessing(c context.Context, req *types.QueryGetBlessingRequest) (*types.QueryGetBlessingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	blessing, found := k.GetBlessing(ctx, req.Index, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBlessingResponse{Blessing: blessing}, nil
}
