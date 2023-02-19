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

func (k Keeper) SignedMessageAll(c context.Context, req *types.QueryAllSignedMessageRequest) (*types.QueryAllSignedMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	msgIdStr := strconv.FormatUint(req.MsgId, 10)

	var signedMessages []types.SignedMessage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signedMessageStore := prefix.NewStore(store, types.KeyPrefix(types.SignedMessageKey+msgIdStr))

	pageRes, err := query.Paginate(signedMessageStore, req.Pagination, func(key []byte, value []byte) error {
		var signedMessage types.SignedMessage
		if err := k.cdc.Unmarshal(value, &signedMessage); err != nil {
			return err
		}

		signedMessages = append(signedMessages, signedMessage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignedMessageResponse{SignedMessage: signedMessages, Pagination: pageRes}, nil
}

func (k Keeper) SignedMessage(c context.Context, req *types.QueryGetSignedMessageRequest) (*types.QueryGetSignedMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	signedMessage, found := k.GetSignedMessage(ctx, req.MsgId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSignedMessageResponse{SignedMessage: signedMessage}, nil
}
