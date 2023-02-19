package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMeditation(goCtx context.Context, msg *types.MsgCreateMeditation) (*types.MsgCreateMeditationResponse, error) {
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

	summmoning, ok := k.GetMagicKeySummoning(ctx, msg.SummoningId)

	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "no summoning exists for this verse")
	}

	expireBlock := summmoning.BlockHeight + 800
	if uint64(ctx.BlockHeight()) > expireBlock {
		// delete this summoning
		k.RemoveMagicKeySummoning(ctx, msg.SummoningId)
		return &types.MsgCreateMeditationResponse{}, nil
	}

	var meditation = types.Meditation{
		Creator:     msg.Creator,
		FromSpirit:  msg.FromSpirit,
		ToSpirit:    msg.ToSpirit,
		WireBytes:   msg.WireBytes,
		Broadcast:   msg.Broadcast,
		ToOld:       msg.ToOld,
		ToOldAndNew: msg.ToOldAndNew,
		MagicKeyId:  msg.MagicKeyId,
		SummoningId: msg.SummoningId,
	}

	id := k.AppendMeditation(
		ctx,
		msg.MagicKeyId,
		meditation,
	)

	return &types.MsgCreateMeditationResponse{
		Id: id,
	}, nil
}
