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

func (k Keeper) SignedScriptureAll(c context.Context, req *types.QueryAllSignedScriptureRequest) (*types.QueryAllSignedScriptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signedScriptures []types.SignedScripture
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signedScriptureStore := prefix.NewStore(store, types.KeyPrefix(types.SignedScriptureKeyPrefix))

	pageRes, err := query.Paginate(signedScriptureStore, req.Pagination, func(key []byte, value []byte) error {
		var signedScripture types.SignedScripture
		if err := k.cdc.Unmarshal(value, &signedScripture); err != nil {
			return err
		}

		signedScriptures = append(signedScriptures, signedScripture)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignedScriptureResponse{SignedScripture: signedScriptures, Pagination: pageRes}, nil
}

func (k Keeper) SignedScripture(c context.Context, req *types.QueryGetSignedScriptureRequest) (*types.QueryGetSignedScriptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSignedScripture(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSignedScriptureResponse{SignedScripture: val}, nil
}
