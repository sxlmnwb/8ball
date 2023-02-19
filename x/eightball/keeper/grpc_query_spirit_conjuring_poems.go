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

func (k Keeper) SpiritConjuringPoemsAll(c context.Context, req *types.QueryAllSpiritConjuringPoemsRequest) (*types.QueryAllSpiritConjuringPoemsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	conjIdStr := strconv.FormatUint(req.MagicId, 10)

	var spiritConjuringPoemss []types.SpiritConjuringPoems
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	spiritConjuringPoemsStore := prefix.NewStore(store, types.KeyPrefix(types.SpiritConjuringPoemsKey+conjIdStr))

	pageRes, err := query.Paginate(spiritConjuringPoemsStore, req.Pagination, func(key []byte, value []byte) error {
		var spiritConjuringPoems types.SpiritConjuringPoems
		if err := k.cdc.Unmarshal(value, &spiritConjuringPoems); err != nil {
			return err
		}

		spiritConjuringPoemss = append(spiritConjuringPoemss, spiritConjuringPoems)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSpiritConjuringPoemsResponse{SpiritConjuringPoems: spiritConjuringPoemss, Pagination: pageRes}, nil
}

func (k Keeper) SpiritConjuringPoems(c context.Context, req *types.QueryGetSpiritConjuringPoemsRequest) (*types.QueryGetSpiritConjuringPoemsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	spiritConjuringPoems, found := k.GetSpiritConjuringPoems(ctx, req.MagicId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSpiritConjuringPoemsResponse{SpiritConjuringPoems: spiritConjuringPoems}, nil
}
