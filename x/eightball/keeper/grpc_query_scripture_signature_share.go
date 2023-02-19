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

func (k Keeper) ScriptureSignatureShareAll(c context.Context, req *types.QueryAllScriptureSignatureShareRequest) (*types.QueryAllScriptureSignatureShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scriptureSignatureShares []types.ScriptureSignatureShare
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	scriptureSignatureShareStore := prefix.NewStore(store, types.KeyPrefix(types.ScriptureSignatureShareKey+req.ScriptureIndex))

	pageRes, err := query.Paginate(scriptureSignatureShareStore, req.Pagination, func(key []byte, value []byte) error {
		var scriptureSignatureShare types.ScriptureSignatureShare
		if err := k.cdc.Unmarshal(value, &scriptureSignatureShare); err != nil {
			return err
		}

		scriptureSignatureShares = append(scriptureSignatureShares, scriptureSignatureShare)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScriptureSignatureShareResponse{ScriptureSignatureShare: scriptureSignatureShares, Pagination: pageRes}, nil
}

func (k Keeper) ScriptureSignatureShare(c context.Context, req *types.QueryGetScriptureSignatureShareRequest) (*types.QueryGetScriptureSignatureShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	scriptureSignatureShare, found := k.GetScriptureSignatureShare(ctx, req.ScriptureIndex, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetScriptureSignatureShareResponse{ScriptureSignatureShare: scriptureSignatureShare}, nil
}
