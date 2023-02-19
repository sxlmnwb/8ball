package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKillMeditationSummoning = "create_kill_meditation_summoning"
	TypeMsgUpdateKillMeditationSummoning = "update_kill_meditation_summoning"
	TypeMsgDeleteKillMeditationSummoning = "delete_kill_meditation_summoning"
)

var _ sdk.Msg = &MsgCreateKillMeditationSummoning{}

func NewMsgCreateKillMeditationSummoning(creator string) *MsgCreateKillMeditationSummoning {
	return &MsgCreateKillMeditationSummoning{
		Creator: creator,
	}
}

func (msg *MsgCreateKillMeditationSummoning) Route() string {
	return RouterKey
}

func (msg *MsgCreateKillMeditationSummoning) Type() string {
	return TypeMsgCreateKillMeditationSummoning
}

func (msg *MsgCreateKillMeditationSummoning) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKillMeditationSummoning) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKillMeditationSummoning) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
