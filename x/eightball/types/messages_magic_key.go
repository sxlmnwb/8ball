package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreateMagicKey = "create_magic_key"
	TypeMsgUpdateMagicKey = "update_magic_key"
)

var _ sdk.Msg = &MsgCreateMagicKey{}

func NewMsgCreateMagicKey(creator string) *MsgCreateMagicKey {
	return &MsgCreateMagicKey{
		Creator: creator,
	}
}

func (msg *MsgCreateMagicKey) Route() string {
	return RouterKey
}

func (msg *MsgCreateMagicKey) Type() string {
	return TypeMsgCreateMagicKey
}

func (msg *MsgCreateMagicKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMagicKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMagicKey) ValidateBasic() error {

	return nil
}
