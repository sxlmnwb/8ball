package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateScripture = "create_scripture"
)

var _ sdk.Msg = &MsgCreateScripture{}

func NewMsgCreateScripture(
	witness string,
	index string,
	hash string,
	location string,
	address string,
	alias string,
	value string,

) *MsgCreateScripture {
	return &MsgCreateScripture{
		Witness:  witness,
		Index:    index,
		Hash:     hash,
		Location: location,
		Address:  address,
		Alias:    alias,
		Value:    value,
	}
}

func (msg *MsgCreateScripture) Route() string {
	return RouterKey
}

func (msg *MsgCreateScripture) Type() string {
	return TypeMsgCreateScripture
}

func (msg *MsgCreateScripture) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateScripture) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateScripture) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
