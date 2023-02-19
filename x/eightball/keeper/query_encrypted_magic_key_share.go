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

func (k Keeper) EncryptedMagicKeyShareAll(c context.Context, req *types.QueryAllEncryptedMagicKeyShareRequest) (*types.QueryAllEncryptedMagicKeyShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var encryptedMagicKeyShares []types.EncryptedMagicKeyShare
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	encryptedMagicKeyShareStore := prefix.NewStore(store, types.KeyPrefix(types.EncryptedMagicKeyShareKeyPrefix))

	pageRes, err := query.Paginate(encryptedMagicKeyShareStore, req.Pagination, func(key []byte, value []byte) error {
		var encryptedMagicKeyShare types.EncryptedMagicKeyShare
		if err := k.cdc.Unmarshal(value, &encryptedMagicKeyShare); err != nil {
			return err
		}

		encryptedMagicKeyShares = append(encryptedMagicKeyShares, encryptedMagicKeyShare)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEncryptedMagicKeyShareResponse{EncryptedMagicKeyShare: encryptedMagicKeyShares, Pagination: pageRes}, nil
}

func (k Keeper) EncryptedMagicKeyShare(c context.Context, req *types.QueryGetEncryptedMagicKeyShareRequest) (*types.QueryGetEncryptedMagicKeyShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEncryptedMagicKeyShare(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEncryptedMagicKeyShareResponse{EncryptedMagicKeyShare: val}, nil
}
