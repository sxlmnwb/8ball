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

func (k Keeper) KillConjuringAll(c context.Context, req *types.QueryAllKillConjuringRequest) (*types.QueryAllKillConjuringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killConjurings []types.KillConjuring
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killConjuringStore := prefix.NewStore(store, types.KeyPrefix(types.KillConjuringKey))

	pageRes, err := query.Paginate(killConjuringStore, req.Pagination, func(key []byte, value []byte) error {
		var killConjuring types.KillConjuring
		if err := k.cdc.Unmarshal(value, &killConjuring); err != nil {
			return err
		}

		killConjurings = append(killConjurings, killConjuring)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillConjuringResponse{KillConjuring: killConjurings, Pagination: pageRes}, nil
}

func (k Keeper) KillConjuring(c context.Context, req *types.QueryGetKillConjuringRequest) (*types.QueryGetKillConjuringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killConjuring, found := k.GetKillConjuring(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillConjuringResponse{KillConjuring: killConjuring}, nil
}
