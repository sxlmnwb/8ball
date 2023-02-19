package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateScriptureSignatureRequest = "create_scripture_signature_request"
	TypeMsgUpdateScriptureSignatureRequest = "update_scripture_signature_request"
	TypeMsgDeleteScriptureSignatureRequest = "delete_scripture_signature_request"
)

var _ sdk.Msg = &MsgCreateScriptureSignatureRequest{}

func NewMsgCreateScriptureSignatureRequest(creator string, scriptureIndex string) *MsgCreateScriptureSignatureRequest {
	return &MsgCreateScriptureSignatureRequest{
		Creator:        creator,
		ScriptureIndex: scriptureIndex,
	}
}

func (msg *MsgCreateScriptureSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateScriptureSignatureRequest) Type() string {
	return TypeMsgCreateScriptureSignatureRequest
}

func (msg *MsgCreateScriptureSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateScriptureSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateScriptureSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
