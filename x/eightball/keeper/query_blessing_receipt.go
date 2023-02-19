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

func (k Keeper) BlessingReceiptAll(c context.Context, req *types.QueryAllBlessingReceiptRequest) (*types.QueryAllBlessingReceiptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blessingReceipts []types.BlessingReceipt
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	blessingReceiptStore := prefix.NewStore(store, types.KeyPrefix(types.BlessingReceiptKeyPrefix))

	pageRes, err := query.Paginate(blessingReceiptStore, req.Pagination, func(key []byte, value []byte) error {
		var blessingReceipt types.BlessingReceipt
		if err := k.cdc.Unmarshal(value, &blessingReceipt); err != nil {
			return err
		}

		blessingReceipts = append(blessingReceipts, blessingReceipt)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlessingReceiptResponse{BlessingReceipt: blessingReceipts, Pagination: pageRes}, nil
}

func (k Keeper) BlessingReceipt(c context.Context, req *types.QueryGetBlessingReceiptRequest) (*types.QueryGetBlessingReceiptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBlessingReceipt(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBlessingReceiptResponse{BlessingReceipt: val}, nil
}
