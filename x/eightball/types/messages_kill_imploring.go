package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillImploring = "create_kill_imploring"
	TypeMsgUpdateKillImploring = "update_kill_imploring"
	TypeMsgDeleteKillImploring = "delete_kill_imploring"
)

var _ sdk.Msg = &MsgCreateKillImploring{}

func NewMsgCreateKillImploring(creator string) *MsgCreateKillImploring {
	return &MsgCreateKillImploring{
		Creator: creator,
	}
}

func (msg *MsgCreateKillImploring) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillImploring) Type() string {
	return TypeMsgCreateKillImploring
}

func (msg *MsgCreateKillImploring) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillImploring) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillImploring) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
