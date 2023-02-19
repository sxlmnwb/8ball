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

func (k Keeper) SignatureRequestAll(c context.Context, req *types.QueryAllSignatureRequestRequest) (*types.QueryAllSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signatureRequests []types.SignatureRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signatureRequestStore := prefix.NewStore(store, types.KeyPrefix(types.SignatureRequestKey))

	pageRes, err := query.Paginate(signatureRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var signatureRequest types.SignatureRequest
		if err := k.cdc.Unmarshal(value, &signatureRequest); err != nil {
			return err
		}

		signatureRequests = append(signatureRequests, signatureRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignatureRequestResponse{SignatureRequest: signatureRequests, Pagination: pageRes}, nil
}

func (k Keeper) SignatureRequest(c context.Context, req *types.QueryGetSignatureRequestRequest) (*types.QueryGetSignatureRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	signatureRequest, found := k.GetSignatureRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSignatureRequestResponse{SignatureRequest: signatureRequest}, nil
}
