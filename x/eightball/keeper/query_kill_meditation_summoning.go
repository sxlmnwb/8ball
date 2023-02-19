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

func (k Keeper) KillMeditationSummoningAll(c context.Context, req *types.QueryAllKillMeditationSummoningRequest) (*types.QueryAllKillMeditationSummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killMeditationSummonings []types.KillMeditationSummoning
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killMeditationSummoningStore := prefix.NewStore(store, types.KeyPrefix(types.KillMeditationSummoningKey))

	pageRes, err := query.Paginate(killMeditationSummoningStore, req.Pagination, func(key []byte, value []byte) error {
		var killMeditationSummoning types.KillMeditationSummoning
		if err := k.cdc.Unmarshal(value, &killMeditationSummoning); err != nil {
			return err
		}

		killMeditationSummonings = append(killMeditationSummonings, killMeditationSummoning)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillMeditationSummoningResponse{KillMeditationSummoning: killMeditationSummonings, Pagination: pageRes}, nil
}

func (k Keeper) KillMeditationSummoning(c context.Context, req *types.QueryGetKillMeditationSummoningRequest) (*types.QueryGetKillMeditationSummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killMeditationSummoning, found := k.GetKillMeditationSummoning(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillMeditationSummoningResponse{KillMeditationSummoning: killMeditationSummoning}, nil
}
