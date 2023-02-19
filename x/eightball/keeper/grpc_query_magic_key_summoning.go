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

func (k Keeper) MagicKeySummoningAll(c context.Context, req *types.QueryAllMagicKeySummoningRequest) (*types.QueryAllMagicKeySummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var magicKeySummonings []types.MagicKeySummoning
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	magicKeySummoningStore := prefix.NewStore(store, types.KeyPrefix(types.MagicKeySummoningKey))

	pageRes, err := query.Paginate(magicKeySummoningStore, req.Pagination, func(key []byte, value []byte) error {
		var magicKeySummoning types.MagicKeySummoning
		if err := k.cdc.Unmarshal(value, &magicKeySummoning); err != nil {
			return err
		}

		magicKeySummonings = append(magicKeySummonings, magicKeySummoning)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMagicKeySummoningResponse{MagicKeySummoning: magicKeySummonings, Pagination: pageRes}, nil
}

func (k Keeper) MagicKeySummoning(c context.Context, req *types.QueryGetMagicKeySummoningRequest) (*types.QueryGetMagicKeySummoningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	magicKeySummoning, found := k.GetMagicKeySummoning(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMagicKeySummoningResponse{MagicKeySummoning: magicKeySummoning}, nil
}
