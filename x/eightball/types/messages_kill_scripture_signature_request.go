package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillScriptureSignatureRequest = "create_kill_scripture_signature_request"
	TypeMsgUpdateKillScriptureSignatureRequest = "update_kill_scripture_signature_request"
	TypeMsgDeleteKillScriptureSignatureRequest = "delete_kill_scripture_signature_request"
)

var _ sdk.Msg = &MsgCreateKillScriptureSignatureRequest{}

func NewMsgCreateKillScriptureSignatureRequest(creator string) *MsgCreateKillScriptureSignatureRequest {
	return &MsgCreateKillScriptureSignatureRequest{
		Creator: creator,
	}
}

func (msg *MsgCreateKillScriptureSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillScriptureSignatureRequest) Type() string {
	return TypeMsgCreateKillScriptureSignatureRequest
}

func (msg *MsgCreateKillScriptureSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillScriptureSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillScriptureSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
