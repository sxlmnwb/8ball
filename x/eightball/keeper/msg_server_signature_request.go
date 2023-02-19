package keeper

import (
	"context"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSignatureRequest(goCtx context.Context, msg *types.MsgCreateSignatureRequest) (*types.MsgCreateSignatureRequestResponse, error) {
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

	var signatureRequest = types.SignatureRequest{
		Creator:   msg.Creator,
		MessageId: msg.MessageId,
	}

	id := k.AppendSignatureRequest(
		ctx,
		signatureRequest,
	)

	return &types.MsgCreateSignatureRequestResponse{
		Id: id,
	}, nil
}
