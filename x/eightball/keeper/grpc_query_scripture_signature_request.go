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

func (k Keeper) ScriptureSignatureRequestAll(c context.Context, req *types.QueryAllScriptureSignatureRequestRequest) (*types.QueryAllScriptureSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scriptureSignatureRequests []types.ScriptureSignatureRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	scriptureSignatureRequestStore := prefix.NewStore(store, types.KeyPrefix(types.ScriptureSignatureRequestKey))

	pageRes, err := query.Paginate(scriptureSignatureRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var scriptureSignatureRequest types.ScriptureSignatureRequest
		if err := k.cdc.Unmarshal(value, &scriptureSignatureRequest); err != nil {
			return err
		}

		scriptureSignatureRequests = append(scriptureSignatureRequests, scriptureSignatureRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScriptureSignatureRequestResponse{ScriptureSignatureRequest: scriptureSignatureRequests, Pagination: pageRes}, nil
}

func (k Keeper) ScriptureSignatureRequest(c context.Context, req *types.QueryGetScriptureSignatureRequestRequest) (*types.QueryGetScriptureSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	scriptureSignatureRequest, found := k.GetScriptureSignatureRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetScriptureSignatureRequestResponse{ScriptureSignatureRequest: scriptureSignatureRequest}, nil
}
