package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVision(goCtx context.Context,
	msg *types.MsgCreateVision) (*types.MsgCreateVisionResponse, error) {
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

	var vision = types.Vision{
		Creator:     msg.Creator,
		MagicKeyId:  msg.MagicKeyId,
		EcPointX:    msg.EcPointX,
		EcPointY:    msg.EcPointY,
		SummoningId: msg.SummoningId,
	}

	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	thisVisionPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)

	if thisVisionPower.Equal(math.NewInt(0)) {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Delegator: %s has no power", msg.Creator))

	}

	// get existing visions for this key
	visions := k.GetAllVision(ctx, msg.MagicKeyId)

	// existing vision power
	existingVisionPower := math.NewInt(0)
	for _, vision := range visions {
		visionCreator := vision.Creator
		creator, err := sdk.AccAddressFromBech32(visionCreator)
		if err != nil {
			// can not find currrent account address from string
			continue
		}
		bondedPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
		if vision.MagicKeyId != msg.MagicKeyId {
			continue
		}

		if vision.EcPointX != msg.EcPointX {
			continue
		}

		if vision.EcPointY != msg.EcPointY {
			continue
		}

		if vision.SummoningId != msg.SummoningId {
			continue
		}

		existingVisionPower = existingVisionPower.Add(bondedPower)
	}
	totalVisionPower := existingVisionPower.Add(thisVisionPower)
	totalVisionPercent := totalVisionPower.Mul(math.NewInt(100))

	percentVoted := totalVisionPercent.Quo(lastTotalPower)

	id := k.AppendVision(
		ctx,
		msg.MagicKeyId,
		vision,
	)

	if percentVoted.LT(math.NewInt(66)) {
		return &types.MsgCreateVisionResponse{
			Id: id,
		}, nil
	}

	// 66% + vision Power voting for this magic key

	xPoint := msg.GetEcPointX()
	yPoint := msg.GetEcPointY()

	validPubKey := checkValidPublicKey(xPoint, yPoint)

	if !validPubKey {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Pub Key X: %x Y: %s \n", xPoint, yPoint))
	}

	bitcoinAddress, ok := getBitcoinAddress(xPoint, yPoint)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Bitcoin Alias for Key X: %x Y: %s \n", xPoint, yPoint))
	}
	ethereumAddress, ok := getEthereumAddress(xPoint, yPoint)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Ethereum Alias for Key X: %x Y: %s \n", xPoint, yPoint))
	}
	tronAddress, ok := getTronAddress(xPoint, yPoint)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Tron Alias for Key X: %x Y: %s \n", xPoint, yPoint))
	}
	zclassicAddress, ok := getZclassicAddress(xPoint, yPoint)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Zclassic Alias for Key X: %x Y: %s \n", xPoint, yPoint))
	}
	eightballAddress, ok := getEightballAddress(xPoint, yPoint)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("Invalid Eightball Alias for Key X: %x Y: %s \n", xPoint, yPoint))
	}

	// Set magic key EC address
	var magicKey = types.MagicKey{
		Id:             msg.MagicKeyId,
		BlockHeight:    uint64(ctx.BlockHeight()),
		EcPointX:       xPoint,
		EcPointY:       yPoint,
		BitcoinAlias:   bitcoinAddress,
		EthereumAlias:  ethereumAddress,
		ZclassicAlias:  zclassicAddress,
		TronAlias:      tronAddress,
		EightballAlias: eightballAddress,
	}
	k.SetMagicKey(ctx, magicKey)

	// delete this summoning
	k.RemoveMagicKeySummoning(ctx, msg.SummoningId)
	// set current magic key
	var currentMagicKey = types.CurrentMagicKey{
		SinceBlock:      uint64(ctx.BlockHeight()),
		CurrentMagicKey: msg.MagicKeyId,
	}

	k.SetCurrentMagicKey(ctx, currentMagicKey)

	return &types.MsgCreateVisionResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteVision(goCtx context.Context,
	msg *types.MsgDeleteVision) (*types.MsgDeleteVisionResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
		"can not delete vision")
}
