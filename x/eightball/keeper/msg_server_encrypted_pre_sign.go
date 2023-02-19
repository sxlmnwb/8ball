package keeper

import (
	"context"

	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateEncryptedPreSign(goCtx context.Context, msg *types.MsgCreateEncryptedPreSign) (*types.MsgCreateEncryptedPreSignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	indexString := msg.Creator + "-magic-key-" + msg.MagicKeyId

	if !isBonded(ctx, k, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"message creator not bonded with at least 8%",
		)
	}

	var encryptedPreSign = types.EncryptedPreSign{
		Creator: msg.Creator,
		Index:   indexString,
		Data:    msg.Data,
	}

	k.SetEncryptedPreSign(
		ctx,
		encryptedPreSign,
	)
	return &types.MsgCreateEncryptedPreSignResponse{}, nil
}

func (k msgServer) DeleteEncryptedPreSign(goCtx context.Context, msg *types.MsgDeleteEncryptedPreSign) (*types.MsgDeleteEncryptedPreSignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	indexString := msg.Creator + "-magic-key-" + msg.MagicKeyId

	// Check if the value exists
	valFound, isFound := k.GetEncryptedPreSign(
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

	k.RemoveEncryptedPreSign(
		ctx,
		indexString,
	)

	return &types.MsgDeleteEncryptedPreSignResponse{}, nil
}
