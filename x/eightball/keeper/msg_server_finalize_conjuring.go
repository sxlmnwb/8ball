package keeper

import (
	b64 "encoding/base64"

	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) FinalizeConjuring(goCtx context.Context,
	msg *types.MsgFinalizeConjuring) (*types.MsgFinalizeConjuringResponse, error) {
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

	// get conjuring object
	hcConj, hccFound := k.GetHighCouncilConjurings(ctx, msg.GetConjuringId())

	if !hccFound {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not find conjuring with id  %d ", msg.GetConjuringId()))

	}

	// kill any expired summmonings
	allSummoning := k.GetAllMagicKeySummoning(ctx)
	for _, summoning := range allSummoning {
		expBlock := summoning.BlockHeight + 800
		if uint64(ctx.BlockHeight()) > expBlock {
			// delete this summoning
			k.RemoveMagicKeySummoning(ctx, summoning.Id)
		}
	}

	// must wait at least 5 blocks after conjuring till finalizing
	readyBlock := hcConj.BlockStarted + 5
	if uint64(ctx.BlockHeight()) < readyBlock {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("conjuring not ready to finalize.  curr block: %d, ready_block %d", ctx.BlockHeight(), readyBlock))
	}

	// if we are 15+ blocks past the ready block, fail this conjuring
	expireBlock := readyBlock + uint64(15)
	if uint64(ctx.BlockHeight()) > expireBlock {
		// delete conjuring object
		k.RemoveHighCouncilConjurings(ctx, hcConj.Id)

		return &types.MsgFinalizeConjuringResponse{}, nil
	}

	// get spirit poems for conjuring
	poems := k.GetAllSpiritConjuringPoems(ctx, msg.ConjuringId)

	// build high council
	// get all unique spirits with poems for this conjuring
	var oratorMap = make(map[string]bool)

	for _, poem := range poems {
		orator := poem.Creator
		sdkAddr, err := sdk.AccAddressFromBech32(orator)
		if err != nil {
			return nil, sdkerrors.Wrap(
				sdkerrors.ErrLogic,
				fmt.Sprintf("can not find address for poem orator  %s ", orator))

		}

		pubKeyBytes := k.accountKeeper.GetAccount(ctx, sdkAddr).GetPubKey().Bytes()
		pubKey := b64.StdEncoding.EncodeToString(pubKeyBytes)

		newSpirit := types.HighCouncil{
			MagicKeyId:    hcConj.MagicKeyId,
			SpiritAddress: orator,
			SpiritPubKey:  pubKey,
		}
		if oratorMap[orator] {
			continue
		}
		k.AppendHighCouncil(ctx, hcConj.MagicKeyId, newSpirit)
		oratorMap[orator] = true
	}

	// generate magic key summoning
	k.AppendMagicKeySummoning(
		ctx,
		types.MagicKeySummoning{
			MagicKeyId:  hcConj.MagicKeyId,
			BlockHeight: uint64(ctx.BlockHeight()),
		},
	)

	// delete conjuring object
	k.RemoveHighCouncilConjurings(ctx, hcConj.Id)

	return &types.MsgFinalizeConjuringResponse{}, nil
}
