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

func (k Keeper) MagicKeyAll(c context.Context, req *types.QueryAllMagicKeyRequest) (*types.QueryAllMagicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var magicKeys []types.MagicKey
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	magicKeyStore := prefix.NewStore(store, types.KeyPrefix(types.MagicKeyKey))

	pageRes, err := query.Paginate(magicKeyStore, req.Pagination, func(key []byte, value []byte) error {
		var magicKey types.MagicKey
		if err := k.cdc.Unmarshal(value, &magicKey); err != nil {
			return err
		}

		magicKeys = append(magicKeys, magicKey)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMagicKeyResponse{MagicKey: magicKeys, Pagination: pageRes}, nil
}

func (k Keeper) MagicKey(c context.Context, req *types.QueryGetMagicKeyRequest) (*types.QueryGetMagicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	magicKey, found := k.GetMagicKey(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMagicKeyResponse{MagicKey: magicKey}, nil
}
