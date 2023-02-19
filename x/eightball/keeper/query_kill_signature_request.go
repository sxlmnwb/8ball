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

func (k Keeper) KillSignatureRequestAll(c context.Context, req *types.QueryAllKillSignatureRequestRequest) (*types.QueryAllKillSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var killSignatureRequests []types.KillSignatureRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	killSignatureRequestStore := prefix.NewStore(store, types.KeyPrefix(types.KillSignatureRequestKey))

	pageRes, err := query.Paginate(killSignatureRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var killSignatureRequest types.KillSignatureRequest
		if err := k.cdc.Unmarshal(value, &killSignatureRequest); err != nil {
			return err
		}

		killSignatureRequests = append(killSignatureRequests, killSignatureRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKillSignatureRequestResponse{KillSignatureRequest: killSignatureRequests, Pagination: pageRes}, nil
}

func (k Keeper) KillSignatureRequest(c context.Context, req *types.QueryGetKillSignatureRequestRequest) (*types.QueryGetKillSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	killSignatureRequest, found := k.GetKillSignatureRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKillSignatureRequestResponse{KillSignatureRequest: killSignatureRequest}, nil
}
