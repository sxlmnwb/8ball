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

func (k Keeper) SignedScriptureListAll(c context.Context, req *types.QueryAllSignedScriptureListRequest) (*types.QueryAllSignedScriptureListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signedScriptureLists []types.SignedScriptureList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signedScriptureListStore := prefix.NewStore(store, types.KeyPrefix(types.SignedScriptureListKey))

	pageRes, err := query.Paginate(signedScriptureListStore, req.Pagination, func(key []byte, value []byte) error {
		var signedScriptureList types.SignedScriptureList
		if err := k.cdc.Unmarshal(value, &signedScriptureList); err != nil {
			return err
		}

		signedScriptureLists = append(signedScriptureLists, signedScriptureList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignedScriptureListResponse{SignedScriptureList: signedScriptureLists, Pagination: pageRes}, nil
}

func (k Keeper) SignedScriptureList(c context.Context, req *types.QueryGetSignedScriptureListRequest) (*types.QueryGetSignedScriptureListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	signedScriptureList, found := k.GetSignedScriptureList(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSignedScriptureListResponse{SignedScriptureList: signedScriptureList}, nil
}
