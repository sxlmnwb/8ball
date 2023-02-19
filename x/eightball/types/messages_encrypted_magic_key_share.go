package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEncryptedMagicKeyShare = "create_encrypted_magic_key_share"
	TypeMsgUpdateEncryptedMagicKeyShare = "update_encrypted_magic_key_share"
	TypeMsgDeleteEncryptedMagicKeyShare = "delete_encrypted_magic_key_share"
)

var _ sdk.Msg = &MsgCreateEncryptedMagicKeyShare{}

func NewMsgCreateEncryptedMagicKeyShare(
	creator string,
	magicKeyId string,
	data string,

) *MsgCreateEncryptedMagicKeyShare {
	return &MsgCreateEncryptedMagicKeyShare{
		Creator:    creator,
		MagicKeyId: magicKeyId,
		Data:       data,
	}
}

func (msg *MsgCreateEncryptedMagicKeyShare) Route() string {
	return RouterKey
}

func (msg *MsgCreateEncryptedMagicKeyShare) Type() string {
	return TypeMsgCreateEncryptedMagicKeyShare
}

func (msg *MsgCreateEncryptedMagicKeyShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEncryptedMagicKeyShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEncryptedMagicKeyShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEncryptedMagicKeyShare{}

func NewMsgDeleteEncryptedMagicKeyShare(
	creator string,
	magicKeyId string,

) *MsgDeleteEncryptedMagicKeyShare {
	return &MsgDeleteEncryptedMagicKeyShare{
		Creator:    creator,
		MagicKeyId: magicKeyId,
	}
}
func (msg *MsgDeleteEncryptedMagicKeyShare) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEncryptedMagicKeyShare) Type() string {
	return TypeMsgDeleteEncryptedMagicKeyShare
}

func (msg *MsgDeleteEncryptedMagicKeyShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEncryptedMagicKeyShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEncryptedMagicKeyShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
