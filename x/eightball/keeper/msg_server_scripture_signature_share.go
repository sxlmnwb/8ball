package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateScriptureSignatureShare(goCtx context.Context, msg *types.MsgCreateScriptureSignatureShare) (*types.MsgCreateScriptureSignatureShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// require request to come from bonded spirit
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Creator))
	}
	thisSpiritPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisSpiritPower = thisSpiritPower.Quo(math.NewInt(1000000))
	if thisSpiritPower.IsNil() || thisSpiritPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Creator))
	}

	ueblPower := k.stakingKeeper.GetDelegatorBonded(ctx,
		sdk.AccAddress(creator))

	if ueblPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Creator))

	}

	var scriptureSignatureShare = types.ScriptureSignatureShare{
		Creator:        msg.Creator,
		ScriptureIndex: msg.ScriptureIndex,
		ShareData:      msg.ShareData,
		MagicKeyId:     msg.MagicKeyId,
		PubKey:         msg.PubKey,
	}
	existingSignatureShare := k.GetAllScriptureSignatureShare(ctx,
		msg.ScriptureIndex)

	for _, v := range existingSignatureShare {
		// overwrite share data from same Creator
		if v.Creator == msg.Creator {
			scriptureSignatureShare.Id = v.Id
			k.SetScriptureSignatureShare(ctx, msg.ScriptureIndex, scriptureSignatureShare)
			return &types.MsgCreateScriptureSignatureShareResponse{
				Id: v.Id,
			}, nil
		}
	}

	id := k.AppendScriptureSignatureShare(
		ctx,
		msg.ScriptureIndex,
		scriptureSignatureShare,
	)

	return &types.MsgCreateScriptureSignatureShareResponse{
		Id: id,
	}, nil
}
