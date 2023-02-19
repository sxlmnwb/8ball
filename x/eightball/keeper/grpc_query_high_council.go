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

func (k Keeper) HighCouncilAll(c context.Context, req *types.QueryAllHighCouncilRequest) (*types.QueryAllHighCouncilResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	keyId := req.GetKeyId()
	keyIdStr := strconv.FormatUint(keyId, 10)

	var highCouncils []types.HighCouncil
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	highCouncilStore := prefix.NewStore(store, types.KeyPrefix(types.HighCouncilKey+keyIdStr))

	pageRes, err := query.Paginate(highCouncilStore, req.Pagination, func(key []byte, value []byte) error {
		var highCouncil types.HighCouncil
		if err := k.cdc.Unmarshal(value, &highCouncil); err != nil {
			return err
		}

		highCouncils = append(highCouncils, highCouncil)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHighCouncilResponse{HighCouncil: highCouncils, Pagination: pageRes}, nil
}

func (k Keeper) HighCouncil(c context.Context, req *types.QueryGetHighCouncilRequest) (*types.QueryGetHighCouncilResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	highCouncil, found := k.GetHighCouncil(ctx, req.KeyId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHighCouncilResponse{HighCouncil: highCouncil}, nil
}
