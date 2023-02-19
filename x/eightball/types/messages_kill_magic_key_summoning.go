package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillMagicKeySummoning = "create_kill_magic_key_summoning"
	TypeMsgUpdateKillMagicKeySummoning = "update_kill_magic_key_summoning"
	TypeMsgDeleteKillMagicKeySummoning = "delete_kill_magic_key_summoning"
)

var _ sdk.Msg = &MsgCreateKillMagicKeySummoning{}

func NewMsgCreateKillMagicKeySummoning(creator string) *MsgCreateKillMagicKeySummoning {
	return &MsgCreateKillMagicKeySummoning{
		Creator: creator,
	}
}

func (msg *MsgCreateKillMagicKeySummoning) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillMagicKeySummoning) Type() string {
	return TypeMsgCreateKillMagicKeySummoning
}

func (msg *MsgCreateKillMagicKeySummoning) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillMagicKeySummoning) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillMagicKeySummoning) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
