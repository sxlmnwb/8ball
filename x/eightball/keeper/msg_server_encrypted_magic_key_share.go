package keeper

import (
	"context"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateEncryptedMagicKeyShare(goCtx context.Context, msg *types.MsgCreateEncryptedMagicKeyShare) (*types.MsgCreateEncryptedMagicKeyShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !isBonded(ctx, k, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"message creator not bonded with at least 8%",
		)
	}

	indexString := msg.Creator + "-magic-key-" + msg.MagicKeyId

	var encryptedMagicKeyShare = types.EncryptedMagicKeyShare{
		Creator: msg.Creator,
		Index:   indexString,
		Data:    msg.Data,
	}

	k.SetEncryptedMagicKeyShare(
		ctx,
		encryptedMagicKeyShare,
	)
	return &types.MsgCreateEncryptedMagicKeyShareResponse{}, nil
}

func (k msgServer) DeleteEncryptedMagicKeyShare(goCtx context.Context, msg *types.MsgDeleteEncryptedMagicKeyShare) (*types.MsgDeleteEncryptedMagicKeyShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !isBonded(ctx, k, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"message creator not bonded with at least 8%",
		)
	}

	indexString := msg.Creator + "-magic-key-" + msg.MagicKeyId

	// Check if the value exists
	valFound, isFound := k.GetEncryptedMagicKeyShare(
		ctx,
		indexString,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEncryptedMagicKeyShare(
		ctx,
		indexString,
	)

	return &types.MsgDeleteEncryptedMagicKeyShareResponse{}, nil
}

func isBonded(ctx sdk.Context, k msgServer, creatorString string) bool {
	// require request to come from bonded spirit
	creator, err := sdk.AccAddressFromBech32(creatorString)
	if err != nil {
		return false
	}
	thisSpiritPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisSpiritPower = thisSpiritPower.Quo(math.NewInt(1000000))
	if thisSpiritPower.IsNil() || thisSpiritPower.IsZero() {
		return false
	}

	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	thisSpiritPercent := thisSpiritPower.Mul(math.NewInt(100))

	percentVoted := thisSpiritPercent.Quo(lastTotalPower)

	// require at least 8% power proceed
	if percentVoted.LT(math.NewInt(8)) {
		return false
	}
	return true
}
