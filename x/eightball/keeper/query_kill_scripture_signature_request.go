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

func (k Keeper) KillScriptureSignatureRequestAll(c context.Context, req *types.QueryAllKillScriptureSignatureRequestRequest) (*types.QueryAllKillScriptureSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killScriptureSignatureRequests []types.KillScriptureSignatureRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killScriptureSignatureRequestStore := prefix.NewStore(store, types.KeyPrefix(types.KillScriptureSignatureRequestKey))

	pageRes, err := query.Paginate(killScriptureSignatureRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var killScriptureSignatureRequest types.KillScriptureSignatureRequest
		if err := k.cdc.Unmarshal(value, &killScriptureSignatureRequest); err != nil {
			return err
		}

		killScriptureSignatureRequests = append(killScriptureSignatureRequests, killScriptureSignatureRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillScriptureSignatureRequestResponse{KillScriptureSignatureRequest: killScriptureSignatureRequests, Pagination: pageRes}, nil
}

func (k Keeper) KillScriptureSignatureRequest(c context.Context, req *types.QueryGetKillScriptureSignatureRequestRequest) (*types.QueryGetKillScriptureSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killScriptureSignatureRequest, found := k.GetKillScriptureSignatureRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillScriptureSignatureRequestResponse{KillScriptureSignatureRequest: killScriptureSignatureRequest}, nil
}
