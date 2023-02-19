package keeper

import (
	"context"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ScriptureAll(c context.Context, req *types.QueryAllScriptureRequest) (*types.QueryAllScriptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scriptures []types.Scripture
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	scriptureStore := prefix.NewStore(store, types.KeyPrefix(types.ScriptureKeyPrefix))

	pageRes, err := query.Paginate(scriptureStore, req.Pagination, func(key []byte, value []byte) error {
		var scripture types.Scripture
		if err := k.cdc.Unmarshal(value, &scripture); err != nil {
			return err
		}

		scriptures = append(scriptures, scripture)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScriptureResponse{Scripture: scriptures, Pagination: pageRes}, nil
}

func (k Keeper) Scripture(c context.Context, req *types.QueryGetScriptureRequest) (*types.QueryGetScriptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetScripture(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetScriptureResponse{Scripture: val}, nil
}
