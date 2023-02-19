package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEncryptedPreSign = "create_encrypted_pre_sign"
	TypeMsgUpdateEncryptedPreSign = "update_encrypted_pre_sign"
	TypeMsgDeleteEncryptedPreSign = "delete_encrypted_pre_sign"
)

var _ sdk.Msg = &MsgCreateEncryptedPreSign{}

func NewMsgCreateEncryptedPreSign(
	creator string,
	magicKeyId string,
	data string,

) *MsgCreateEncryptedPreSign {
	return &MsgCreateEncryptedPreSign{
		Creator:    creator,
		MagicKeyId: magicKeyId,
		Data:       data,
	}
}

func (msg *MsgCreateEncryptedPreSign) Route() string {
	return RouterKey
}

func (msg *MsgCreateEncryptedPreSign) Type() string {
	return TypeMsgCreateEncryptedPreSign
}

func (msg *MsgCreateEncryptedPreSign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEncryptedPreSign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEncryptedPreSign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEncryptedPreSign{}

func NewMsgDeleteEncryptedPreSign(
	creator string,
	magicKeyId string,

) *MsgDeleteEncryptedPreSign {
	return &MsgDeleteEncryptedPreSign{
		Creator:    creator,
		MagicKeyId: magicKeyId,
	}
}
func (msg *MsgDeleteEncryptedPreSign) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEncryptedPreSign) Type() string {
	return TypeMsgDeleteEncryptedPreSign
}

func (msg *MsgDeleteEncryptedPreSign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEncryptedPreSign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEncryptedPreSign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
