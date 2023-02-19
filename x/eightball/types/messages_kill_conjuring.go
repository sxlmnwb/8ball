package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillConjuring = "create_kill_conjuring"
	TypeMsgUpdateKillConjuring = "update_kill_conjuring"
	TypeMsgDeleteKillConjuring = "delete_kill_conjuring"
)

var _ sdk.Msg = &MsgCreateKillConjuring{}

func NewMsgCreateKillConjuring(creator string) *MsgCreateKillConjuring {
	return &MsgCreateKillConjuring{
		Creator: creator,
	}
}

func (msg *MsgCreateKillConjuring) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillConjuring) Type() string {
	return TypeMsgCreateKillConjuring
}

func (msg *MsgCreateKillConjuring) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillConjuring) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillConjuring) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
