package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSignatureShare(goCtx context.Context, msg *types.MsgCreateSignatureShare) (*types.MsgCreateSignatureShareResponse, error) {
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

	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	thisSpiritPercent := thisSpiritPower.Mul(math.NewInt(100))

	percentVoted := thisSpiritPercent.Quo(lastTotalPower)

	// require at least 8% power proceed
	if percentVoted.LT(math.NewInt(8)) {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("not enough power to submit a poem thisSpiritPower %s, lastTotalPower %s", thisSpiritPercent, lastTotalPower))
	}

	var signatureShare = types.SignatureShare{
		Creator:    msg.Creator,
		MessageId:  msg.MessageId,
		ShareData:  msg.ShareData,
		MagicKeyId: msg.MagicKeyId,
		PubKey:     msg.PubKey,
	}

	existingSignatureShare := k.GetAllSignatureShare(ctx, msg.MessageId)

	for _, v := range existingSignatureShare {
		// overwrite share data from same PubKey
		if v.PubKey == msg.PubKey {
			k.SetSignatureShare(ctx, msg.MessageId, signatureShare)
			return &types.MsgCreateSignatureShareResponse{
				Id: msg.MessageId,
			}, nil

		}
	}

	id := k.AppendSignatureShare(
		ctx,
		msg.MessageId,
		signatureShare,
	)

	return &types.MsgCreateSignatureShareResponse{
		Id: id,
	}, nil
}
