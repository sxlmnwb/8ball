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

func (k Keeper) HighCouncilConjuringsAll(c context.Context, req *types.QueryAllHighCouncilConjuringsRequest) (*types.QueryAllHighCouncilConjuringsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var highCouncilConjuringss []types.HighCouncilConjurings
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	highCouncilConjuringsStore := prefix.NewStore(store, types.KeyPrefix(types.HighCouncilConjuringsKey))

	pageRes, err := query.Paginate(highCouncilConjuringsStore, req.Pagination, func(key []byte, value []byte) error {
		var highCouncilConjurings types.HighCouncilConjurings
		if err := k.cdc.Unmarshal(value, &highCouncilConjurings); err != nil {
			return err
		}

		highCouncilConjuringss = append(highCouncilConjuringss, highCouncilConjurings)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHighCouncilConjuringsResponse{HighCouncilConjurings: highCouncilConjuringss, Pagination: pageRes}, nil
}

func (k Keeper) HighCouncilConjurings(c context.Context, req *types.QueryGetHighCouncilConjuringsRequest) (*types.QueryGetHighCouncilConjuringsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	highCouncilConjurings, found := k.GetHighCouncilConjurings(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHighCouncilConjuringsResponse{HighCouncilConjurings: highCouncilConjurings}, nil
}
