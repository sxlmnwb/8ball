package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSpiritConjuringPoems(goCtx context.Context, msg *types.MsgCreateSpiritConjuringPoems) (*types.MsgCreateSpiritConjuringPoemsResponse, error) {
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

	// require at least 8% power to submit a poem
	if percentVoted.LT(math.NewInt(8)) {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("not enough power to submit a poem thisSpiritPower %s, lastTotalPower %s", thisSpiritPercent, lastTotalPower))
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

	// kill any expired scripture Signature Requests
	allScripSigRequests := k.GetAllScriptureSignatureRequest(ctx)
	for _, scripSigRequest := range allScripSigRequests {
		expBlock := scripSigRequest.BlockHeight + 800
		if uint64(ctx.BlockHeight()) > expBlock {
			// delete this summoning
			k.RemoveScriptureSignatureRequest(ctx, scripSigRequest.Id)
		}
	}

	// kill any expired scripture Signature Requests
	allSigRequests := k.GetAllSignatureRequest(ctx)
	for _, sigRequest := range allSigRequests {
		expBlock := sigRequest.BlockHeight + 800
		if uint64(ctx.BlockHeight()) > expBlock {
			// delete this summoning
			k.RemoveSignatureRequest(ctx, sigRequest.Id)
		}
	}

	hcc, hccFound := k.GetHighCouncilConjurings(ctx, msg.ConjuringId)
	if !hccFound {
		if !hccFound {
			return nil, sdkerrors.Wrap(
				sdkerrors.ErrKeyNotFound,
				fmt.Sprintf("can not find high council conjuring with id  %d ", msg.ConjuringId))
		}
	}
	readyBlock := hcc.BlockStarted + 5
	// if we are 15+ blocks past the ready block, fail this conjuring
	expireBlock := readyBlock + uint64(15)
	if uint64(ctx.BlockHeight()) > expireBlock {
		// delete conjuring object
		k.RemoveHighCouncilConjurings(ctx, hcc.Id)

		return &types.MsgCreateSpiritConjuringPoemsResponse{}, nil
	}

	magicKey, mkFound := k.GetMagicKey(ctx, hcc.MagicKeyId)
	if !mkFound {
		if !hccFound {
			return nil, sdkerrors.Wrap(
				sdkerrors.ErrKeyNotFound,
				fmt.Sprintf("can not find magic key with id  %d ", hcc.MagicKeyId))
		}
	}

	ueblPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)

	var spiritConjuringPoems = types.SpiritConjuringPoems{
		Creator:     msg.Creator,
		MagicKeyId:  magicKey.Id,
		BlockHeight: uint64(ctx.BlockHeight()),
		Poem:        msg.Poem,
		ConjuringId: msg.ConjuringId,
		UeblPower:   ueblPower.Uint64(),
	}

	id := k.AppendSpiritConjuringPoems(
		ctx,
		magicKey.Id,
		spiritConjuringPoems,
	)

	return &types.MsgCreateSpiritConjuringPoemsResponse{
		Id: id,
	}, nil
}
