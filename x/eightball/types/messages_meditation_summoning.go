package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMeditationSummoning = "create_meditation_summoning"
	TypeMsgUpdateMeditationSummoning = "update_meditation_summoning"
	TypeMsgDeleteMeditationSummoning = "delete_meditation_summoning"
)

var _ sdk.Msg = &MsgCreateMeditationSummoning{}

func NewMsgCreateMeditationSummoning(creator string, magicKeyId uint64) *MsgCreateMeditationSummoning {
	return &MsgCreateMeditationSummoning{
		Creator:    creator,
		MagicKeyId: magicKeyId,
	}
}

func (msg *MsgCreateMeditationSummoning) Route() string {
	return RouterKey
}

func (msg *MsgCreateMeditationSummoning) Type() string {
	return TypeMsgCreateMeditationSummoning
}

func (msg *MsgCreateMeditationSummoning) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMeditationSummoning) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMeditationSummoning) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMeditationSummoning{}

func NewMsgDeleteMeditationSummoning(creator string, id uint64) *MsgDeleteMeditationSummoning {
	return &MsgDeleteMeditationSummoning{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMeditationSummoning) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMeditationSummoning) Type() string {
	return TypeMsgDeleteMeditationSummoning
}

func (msg *MsgDeleteMeditationSummoning) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMeditationSummoning) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMeditationSummoning) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
