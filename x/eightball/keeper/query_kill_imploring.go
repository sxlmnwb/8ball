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

func (k Keeper) KillImploringAll(c context.Context, req *types.QueryAllKillImploringRequest) (*types.QueryAllKillImploringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killImplorings []types.KillImploring
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killImploringStore := prefix.NewStore(store, types.KeyPrefix(types.KillImploringKey))

	pageRes, err := query.Paginate(killImploringStore, req.Pagination, func(key []byte, value []byte) error {
		var killImploring types.KillImploring
		if err := k.cdc.Unmarshal(value, &killImploring); err != nil {
			return err
		}

		killImplorings = append(killImplorings, killImploring)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillImploringResponse{KillImploring: killImplorings, Pagination: pageRes}, nil
}

func (k Keeper) KillImploring(c context.Context, req *types.QueryGetKillImploringRequest) (*types.QueryGetKillImploringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killImploring, found := k.GetKillImploring(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillImploringResponse{KillImploring: killImploring}, nil
}
