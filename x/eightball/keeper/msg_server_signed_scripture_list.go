package keeper

import (
	"context"

	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSignedScriptureList(goCtx context.Context, msg *types.MsgCreateSignedScriptureList) (*types.MsgCreateSignedScriptureListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var signedScriptureList = types.SignedScriptureList{
		Creator:        msg.Creator,
		ScriptureIndex: msg.ScriptureIndex,
	}

	id := k.AppendSignedScriptureList(
		ctx,
		signedScriptureList,
	)

	return &types.MsgCreateSignedScriptureListResponse{
		Id: id,
	}, nil
}
