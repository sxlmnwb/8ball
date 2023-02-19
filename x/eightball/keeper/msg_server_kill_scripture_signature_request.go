package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKillScriptureSignatureRequest(goCtx context.Context, msg *types.MsgCreateKillScriptureSignatureRequest) (*types.MsgCreateKillScriptureSignatureRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Creator))
	}

	thisKillingPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisKillingPower = thisKillingPower.Quo(math.NewInt(1000000))

	if thisKillingPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Creator))

	}

	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)

	// get existing Killings for this key
	existingKillings := k.GetAllKillScriptureSignatureRequest(ctx)

	// existing Killing power
	existingKillingPower := math.NewInt(0)
	for _, existingKilling := range existingKillings {
		existingKillingCreator := existingKilling.Creator
		creator, err := sdk.AccAddressFromBech32(existingKillingCreator)
		if err != nil {
			// can not find currrent account address from string
			continue
		}
		bondedPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
		bondedPower = bondedPower.Quo(math.NewInt(1000000))
		existingKillingPower = existingKillingPower.Add(bondedPower)
	}
	totalKillingPower := existingKillingPower.Add(thisKillingPower)

	totalKillingPercent := totalKillingPower.Mul(math.NewInt(100))

	metadata := fmt.Sprintf("thisKillingPower %s, totalKillingPower %s, lastTotalPower %s creator %s", thisKillingPower, totalKillingPower, lastTotalPower, creator)

	var killScriptureSignatureRequest = types.KillScriptureSignatureRequest{
		Creator:  msg.Creator,
		Height:   uint64(ctx.BlockHeight()),
		Metadata: metadata,
	}

	percentVoted := totalKillingPercent.Quo(lastTotalPower)

	// reset the state
	if percentVoted.GT(math.NewInt(66)) {
		for _, existingKilling := range existingKillings {
			k.RemoveKillScriptureSignatureRequest(ctx, existingKilling.Id)
		}
		scriptureSignatureRequests := k.GetAllScriptureSignatureRequest(ctx)

		for _, scriptureSignatureRequest := range scriptureSignatureRequests {
			k.RemoveScriptureSignatureRequest(ctx, scriptureSignatureRequest.Id)
		}
		return &types.MsgCreateKillScriptureSignatureRequestResponse{}, nil
	}

	id := k.AppendKillScriptureSignatureRequest(
		ctx,
		killScriptureSignatureRequest,
	)

	return &types.MsgCreateKillScriptureSignatureRequestResponse{
		Id: id,
	}, nil
}
