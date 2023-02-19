package keeper

import (
	"context"
	"fmt"
	"math/big"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBlessing(goCtx context.Context, msg *types.MsgCreateBlessing) (*types.MsgCreateBlessingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetSignedScripture(ctx, msg.Index)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not find signed scripture with index %s \n",
				msg.Index))
	}
	scripture, found := k.GetScripture(ctx, msg.Index)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not find scripture with index %s \n",
				msg.Index))
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Creator))
	}

	thisBlessingPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisBlessingPower = thisBlessingPower.Quo(math.NewInt(1000000))

	if thisBlessingPower.IsNil() || thisBlessingPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Creator))
	}

	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)

	// get existing Blessings for this key
	existingBlessings := k.GetAllBlessing(ctx, msg.Index)

	// existing Blessing power
	existingBlessingPower := math.NewInt(0)
	for _, existingBlessing := range existingBlessings {
		existingBlessingCreator := existingBlessing.Creator
		creator, err := sdk.AccAddressFromBech32(existingBlessingCreator)
		if err != nil {
			// can not find currrent account address from string
			continue
		}
		bondedPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
		bondedPower = bondedPower.Quo(math.NewInt(1000000))
		if existingBlessing.Index != msg.Index {
			continue
		}

		existingBlessingPower = existingBlessingPower.Add(bondedPower)
	}
	totalBlessingPower := existingBlessingPower.Add(thisBlessingPower)

	totalBlessingPercent := totalBlessingPower.Mul(math.NewInt(100))

	percentVoted := totalBlessingPercent.Quo(lastTotalPower)

	metadata := fmt.Sprintf("thisBlessingPower %s, totalBlessingPower %s, lastTotalPower %s creator %s", thisBlessingPower, totalBlessingPower, lastTotalPower, creator)

	var blessing = types.Blessing{
		Creator:     msg.Creator,
		Index:       msg.Index,
		BlockHeight: uint64(ctx.BlockHeight()),
		Metadata:    metadata,
	}

	id := k.AppendBlessing(
		ctx,
		msg.Index,
		blessing,
	)

	debug := fmt.Sprintf("thisBlessingPower: %s, creator %s, totalBlessingPower: %s, lastTotalPower: %s", thisBlessingPower.String(), msg.Creator, totalBlessingPower.String(), lastTotalPower.String())

	// stop execution if we already processed this signed scripture
	_, found = k.GetBlessingReceipt(ctx, msg.Index)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("recipt already found for this signed scripture %s \n%s\n", msg.Index, debug))
	}

	if percentVoted.LT(math.NewInt(70)) {
		return &types.MsgCreateBlessingResponse{
			Id: id,
		}, nil
	}

	// 66%+ blesses this signed scripture and there was never a blessing
	// recipt created yet

	amountInt, ok := math.NewIntFromString(scripture.Value)
	// eth / tron use 18 decimal places,
	// zclassic 'zatoshi' uses 8, eightball ueball is 6
	ethConversion := new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)
	tronConversion := new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)
	zclassicConversion := new(big.Int).Exp(big.NewInt(10), big.NewInt(1), nil)
	convertedAmount := new(big.Int)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not convert amount to int %s \n", scripture.Value))
	}
	eballAccount, err := sdk.AccAddressFromBech32(scripture.Alias)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get account for: %s \n", scripture.Alias))
	}

	if scripture.Location == "ethereum" {
		convertedAmount = new(big.Int).Div((*big.Int)(amountInt.BigInt()),
			ethConversion)
	}

	if scripture.Location == "zclassic" {
		convertedAmount = new(big.Int).Div((*big.Int)(amountInt.BigInt()),
			zclassicConversion)
	}

	if scripture.Location == "tron" {
		convertedAmount = new(big.Int).Div((*big.Int)(amountInt.BigInt()),
			tronConversion)
	}
	blessings, ok := math.NewIntFromString(convertedAmount.String())

	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get correct amount: %s \n",
				convertedAmount.String()))
	}

	if blessings.BigInt().Cmp(big.NewInt(0)) > 0 {
		err = k.MintAndSendToAccount(ctx, eballAccount,
			sdk.Coin{Denom: "uebl", Amount: blessings})
		if err != nil {
			message := fmt.Sprintf(
				"amount %s, convertedAmount %s, blessings %s",
				amountInt.String(), convertedAmount.String(),
				blessings.String())
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
				message+err.Error())
		}
	}

	implorings := k.GetAllImploring(ctx)

	for _, imploring := range implorings {
		if imploring.Index == msg.Index {
			k.RemoveImploring(ctx, imploring.Id)
		}

		// also remove stale implorings
		expireHeight := imploring.Height + 800

		if uint64(ctx.BlockHeight()) > expireHeight {
			k.RemoveImploring(ctx, imploring.Id)
		}
	}

	blessingRecipt := types.BlessingReceipt{
		Index:  msg.Index,
		Height: uint64(ctx.BlockHeight()),
	}
	k.SetBlessingReceipt(ctx, blessingRecipt)

	return &types.MsgCreateBlessingResponse{
		Id: id,
	}, nil
}
