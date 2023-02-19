package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillSignatureRequest = "create_kill_signature_request"
	TypeMsgUpdateKillSignatureRequest = "update_kill_signature_request"
	TypeMsgDeleteKillSignatureRequest = "delete_kill_signature_request"
)

var _ sdk.Msg = &MsgCreateKillSignatureRequest{}

func NewMsgCreateKillSignatureRequest(creator string) *MsgCreateKillSignatureRequest {
	return &MsgCreateKillSignatureRequest{
		Creator: creator,
	}
}

func (msg *MsgCreateKillSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillSignatureRequest) Type() string {
	return TypeMsgCreateKillSignatureRequest
}

func (msg *MsgCreateKillSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
