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

func (k Keeper) SignatureShareAll(c context.Context, req *types.QueryAllSignatureShareRequest) (*types.QueryAllSignatureShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	messageIdStr := strconv.FormatUint(req.MessageId, 10)

	var signatureShares []types.SignatureShare
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signatureShareStore := prefix.NewStore(store, types.KeyPrefix(types.SignatureShareKey+messageIdStr))

	pageRes, err := query.Paginate(signatureShareStore, req.Pagination, func(key []byte, value []byte) error {
		var signatureShare types.SignatureShare
		if err := k.cdc.Unmarshal(value, &signatureShare); err != nil {
			return err
		}

		signatureShares = append(signatureShares, signatureShare)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignatureShareResponse{SignatureShare: signatureShares, Pagination: pageRes}, nil
}

func (k Keeper) SignatureShare(c context.Context, req *types.QueryGetSignatureShareRequest) (*types.QueryGetSignatureShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	signatureShare, found := k.GetSignatureShare(ctx, req.MessageId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSignatureShareResponse{SignatureShare: signatureShare}, nil
}
