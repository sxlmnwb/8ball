package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignedScripture = "create_signed_scripture"
	TypeMsgUpdateSignedScripture = "update_signed_scripture"
	TypeMsgDeleteSignedScripture = "delete_signed_scripture"
)

var _ sdk.Msg = &MsgCreateSignedScripture{}

func NewMsgCreateSignedScripture(
	witness string,
	index string,
	signature string,
	btcAddress string,
	magicKeyId string,
	sigRequestId uint64,

) *MsgCreateSignedScripture {
	return &MsgCreateSignedScripture{
		Witness:      witness,
		Index:        index,
		Signature:    signature,
		BtcAddress:   btcAddress,
		MagicKeyId:   magicKeyId,
		SigRequestId: sigRequestId,
	}
}

func (msg *MsgCreateSignedScripture) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignedScripture) Type() string {
	return TypeMsgCreateSignedScripture
}

func (msg *MsgCreateSignedScripture) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignedScripture) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignedScripture) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
