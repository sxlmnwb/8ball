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

func (k Keeper) KillMagicKeySummoningAll(c context.Context, req *types.QueryAllKillMagicKeySummoningRequest) (*types.QueryAllKillMagicKeySummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killMagicKeySummonings []types.KillMagicKeySummoning
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killMagicKeySummoningStore := prefix.NewStore(store, types.KeyPrefix(types.KillMagicKeySummoningKey))

	pageRes, err := query.Paginate(killMagicKeySummoningStore, req.Pagination, func(key []byte, value []byte) error {
		var killMagicKeySummoning types.KillMagicKeySummoning
		if err := k.cdc.Unmarshal(value, &killMagicKeySummoning); err != nil {
			return err
		}

		killMagicKeySummonings = append(killMagicKeySummonings, killMagicKeySummoning)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillMagicKeySummoningResponse{KillMagicKeySummoning: killMagicKeySummonings, Pagination: pageRes}, nil
}

func (k Keeper) KillMagicKeySummoning(c context.Context, req *types.QueryGetKillMagicKeySummoningRequest) (*types.QueryGetKillMagicKeySummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killMagicKeySummoning, found := k.GetKillMagicKeySummoning(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillMagicKeySummoningResponse{KillMagicKeySummoning: killMagicKeySummoning}, nil
}
