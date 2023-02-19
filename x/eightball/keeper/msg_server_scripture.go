package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateScripture(goCtx context.Context, msg *types.MsgCreateScripture) (*types.MsgCreateScriptureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// here is a spot where you can explicitely exlude txn hashes
	// from being included as scriputure.
	//
	// why might you want to do this?  Perhapse if a txn was included
	// in the genesis file
	excludeGenesisTxns := []string{
		"0x60b1dfbf4ec3ccda2ccba262f3048a411f1e6f38be659487f4c1b3bb8a469fdc",
		"0x3fefa8b401c60de0960daff6338a52cdc17cf73f1654c3751476b3a8e384c39c",
		"0x610b8bd33804f7d11b7752eae6b25e54d1d2d52b685c29da053d3ce6afc32a31",
	}

	for _, txn := range excludeGenesisTxns {
		if txn == msg.Hash {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
				fmt.Sprintf(
					"txn hash excluded from scripture %s \n",
					msg.Hash))

		}
	}

	// require request to come from bonded spirit
	creator, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Witness))
	}
	thisSpiritPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisSpiritPower = thisSpiritPower.Quo(math.NewInt(1000000))
	if thisSpiritPower.IsNil() || thisSpiritPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Witness))
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

	// Check if the value already exists
	_, isFound := k.GetScripture(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var scripture = types.Scripture{
		Index:    msg.Index,
		Hash:     msg.Hash,
		Location: msg.Location,
		Address:  msg.Address,
		Alias:    msg.Alias,
		Value:    msg.Value,
	}

	k.SetScripture(
		ctx,
		scripture,
	)

	var sigRequesst = types.ScriptureSignatureRequest{
		ScriptureIndex: msg.Index,
		BlockHeight:    uint64(ctx.BlockHeight()),
	}

	k.AppendScriptureSignatureRequest(
		ctx,
		sigRequesst,
	)

	return &types.MsgCreateScriptureResponse{}, nil
}
