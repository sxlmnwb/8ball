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

func (k Keeper) MeditationSummoningAll(c context.Context, req *types.QueryAllMeditationSummoningRequest) (*types.QueryAllMeditationSummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var meditationSummonings []types.MeditationSummoning
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	meditationSummoningStore := prefix.NewStore(store, types.KeyPrefix(types.MeditationSummoningKey))

	pageRes, err := query.Paginate(meditationSummoningStore, req.Pagination, func(key []byte, value []byte) error {
		var meditationSummoning types.MeditationSummoning
		if err := k.cdc.Unmarshal(value, &meditationSummoning); err != nil {
			return err
		}

		meditationSummonings = append(meditationSummonings, meditationSummoning)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMeditationSummoningResponse{MeditationSummoning: meditationSummonings, Pagination: pageRes}, nil
}

func (k Keeper) MeditationSummoning(c context.Context, req *types.QueryGetMeditationSummoningRequest) (*types.QueryGetMeditationSummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	meditationSummoning, found := k.GetMeditationSummoning(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMeditationSummoningResponse{MeditationSummoning: meditationSummoning}, nil
}
