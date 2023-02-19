package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBlessing = "create_blessing"
	TypeMsgUpdateBlessing = "update_blessing"
	TypeMsgDeleteBlessing = "delete_blessing"
)

var _ sdk.Msg = &MsgCreateBlessing{}

func NewMsgCreateBlessing(creator string, index string) *MsgCreateBlessing {
	return &MsgCreateBlessing{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateBlessing) Route() string {
	return RouterKey
}

func (msg *MsgCreateBlessing) Type() string {
	return TypeMsgCreateBlessing
}

func (msg *MsgCreateBlessing) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBlessing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBlessing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
