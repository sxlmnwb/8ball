package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignedMessage = "create_signed_message"
	TypeMsgUpdateSignedMessage = "update_signed_message"
	TypeMsgDeleteSignedMessage = "delete_signed_message"
)

var _ sdk.Msg = &MsgCreateSignedMessage{}

func NewMsgCreateSignedMessage(
	creator string,
	body string,
	signature string,
	bitcoinAddress string,
	messageId uint64,
	sigRequestId uint64,
	magicKeyId uint64,
) *MsgCreateSignedMessage {
	return &MsgCreateSignedMessage{
		Creator:        creator,
		Body:           body,
		Signature:      signature,
		BitcoinAddress: bitcoinAddress,
		MessageId:      messageId,
		SigRequestId:   sigRequestId,
		MagicKeyId:     magicKeyId,
	}
}

func (msg *MsgCreateSignedMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignedMessage) Type() string {
	return TypeMsgCreateSignedMessage
}

func (msg *MsgCreateSignedMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignedMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignedMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
