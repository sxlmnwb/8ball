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

func (k Keeper) MeditationAll(c context.Context, req *types.QueryAllMeditationRequest) (*types.QueryAllMeditationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	magKeyIdStr := strconv.FormatUint(req.MagicKeyId, 10)

	var meditations []types.Meditation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	meditationStore := prefix.NewStore(store, types.KeyPrefix(types.MeditationKey+magKeyIdStr))

	pageRes, err := query.Paginate(meditationStore, req.Pagination, func(key []byte, value []byte) error {
		var meditation types.Meditation
		if err := k.cdc.Unmarshal(value, &meditation); err != nil {
			return err
		}

		meditations = append(meditations, meditation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMeditationResponse{Meditation: meditations, Pagination: pageRes}, nil
}

func (k Keeper) Meditation(c context.Context, req *types.QueryGetMeditationRequest) (*types.QueryGetMeditationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	meditation, found := k.GetMeditation(ctx, req.MagicKeyId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMeditationResponse{Meditation: meditation}, nil
}
