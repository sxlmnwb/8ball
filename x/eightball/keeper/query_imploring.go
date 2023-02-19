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

func (k Keeper) ImploringAll(c context.Context, req *types.QueryAllImploringRequest) (*types.QueryAllImploringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var implorings []types.Imploring
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	imploringStore := prefix.NewStore(store, types.KeyPrefix(types.ImploringKey))

	pageRes, err := query.Paginate(imploringStore, req.Pagination, func(key []byte, value []byte) error {
		var imploring types.Imploring
		if err := k.cdc.Unmarshal(value, &imploring); err != nil {
			return err
		}

		implorings = append(implorings, imploring)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllImploringResponse{Imploring: implorings, Pagination: pageRes}, nil
}

func (k Keeper) Imploring(c context.Context, req *types.QueryGetImploringRequest) (*types.QueryGetImploringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	imploring, found := k.GetImploring(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetImploringResponse{Imploring: imploring}, nil
}
