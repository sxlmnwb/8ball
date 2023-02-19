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

func (k Keeper) EncryptedPreSignAll(c context.Context, req *types.QueryAllEncryptedPreSignRequest) (*types.QueryAllEncryptedPreSignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var encryptedPreSigns []types.EncryptedPreSign
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	encryptedPreSignStore := prefix.NewStore(store, types.KeyPrefix(types.EncryptedPreSignKeyPrefix))

	pageRes, err := query.Paginate(encryptedPreSignStore, req.Pagination, func(key []byte, value []byte) error {
		var encryptedPreSign types.EncryptedPreSign
		if err := k.cdc.Unmarshal(value, &encryptedPreSign); err != nil {
			return err
		}

		encryptedPreSigns = append(encryptedPreSigns, encryptedPreSign)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEncryptedPreSignResponse{EncryptedPreSign: encryptedPreSigns, Pagination: pageRes}, nil
}

func (k Keeper) EncryptedPreSign(c context.Context, req *types.QueryGetEncryptedPreSignRequest) (*types.QueryGetEncryptedPreSignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEncryptedPreSign(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEncryptedPreSignResponse{EncryptedPreSign: val}, nil
}
