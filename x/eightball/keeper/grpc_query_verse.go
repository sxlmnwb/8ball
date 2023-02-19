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

func (k Keeper) VerseAll(c context.Context, req *types.QueryAllVerseRequest) (*types.QueryAllVerseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	magKeyIdStr := strconv.FormatUint(req.MagicKeyId, 10)

	var verses []types.Verse
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	verseStore := prefix.NewStore(store, types.KeyPrefix(types.VerseKey+magKeyIdStr))

	pageRes, err := query.Paginate(verseStore, req.Pagination, func(key []byte, value []byte) error {
		var verse types.Verse
		if err := k.cdc.Unmarshal(value, &verse); err != nil {
			return err
		}

		verses = append(verses, verse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVerseResponse{Verse: verses, Pagination: pageRes}, nil
}

func (k Keeper) Verse(c context.Context, req *types.QueryGetVerseRequest) (*types.QueryGetVerseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	verse, found := k.GetVerse(ctx, req.MagicKeyId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVerseResponse{Verse: verse}, nil
}
