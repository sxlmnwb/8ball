package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignatureRequest = "create_signature_request"
	TypeMsgUpdateSignatureRequest = "update_signature_request"
	TypeMsgDeleteSignatureRequest = "delete_signature_request"
)

var _ sdk.Msg = &MsgCreateSignatureRequest{}

func NewMsgCreateSignatureRequest(creator string, messageId uint64) *MsgCreateSignatureRequest {
	return &MsgCreateSignatureRequest{
		Creator:   creator,
		MessageId: messageId,
	}
}

func (msg *MsgCreateSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignatureRequest) Type() string {
	return TypeMsgCreateSignatureRequest
}

func (msg *MsgCreateSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSignatureRequest{}

func NewMsgUpdateSignatureRequest(creator string, id uint64, messageId uint64) *MsgUpdateSignatureRequest {
	return &MsgUpdateSignatureRequest{
		Id:        id,
		Creator:   creator,
		MessageId: messageId,
	}
}

func (msg *MsgUpdateSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSignatureRequest) Type() string {
	return TypeMsgUpdateSignatureRequest
}

func (msg *MsgUpdateSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSignatureRequest{}

func NewMsgDeleteSignatureRequest(creator string, id uint64) *MsgDeleteSignatureRequest {
	return &MsgDeleteSignatureRequest{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSignatureRequest) Type() string {
	return TypeMsgDeleteSignatureRequest
}

func (msg *MsgDeleteSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
