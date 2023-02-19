package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignatureShare = "create_signature_share"
	TypeMsgUpdateSignatureShare = "update_signature_share"
	TypeMsgDeleteSignatureShare = "delete_signature_share"
)

var _ sdk.Msg = &MsgCreateSignatureShare{}

func NewMsgCreateSignatureShare(creator string, messageId uint64, shareData string, magicKeyId uint64, pubKeyString string) *MsgCreateSignatureShare {
	return &MsgCreateSignatureShare{
		Creator:    creator,
		MessageId:  messageId,
		ShareData:  shareData,
		MagicKeyId: magicKeyId,
		PubKey:     pubKeyString,
	}
}

func (msg *MsgCreateSignatureShare) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignatureShare) Type() string {
	return TypeMsgCreateSignatureShare
}

func (msg *MsgCreateSignatureShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignatureShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignatureShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
