package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMeditation = "create_meditation"
	TypeMsgUpdateMeditation = "update_meditation"
	TypeMsgDeleteMeditation = "delete_meditation"
)

var _ sdk.Msg = &MsgCreateMeditation{}

func NewMsgCreateMeditation(creator string, fromSpirit string, toSpirit []string, wireBytes string, broadcast bool, toOld bool, toOldAndNew bool, magicKeyId uint64, summoningId uint64) *MsgCreateMeditation {
	return &MsgCreateMeditation{
		Creator:     creator,
		FromSpirit:  fromSpirit,
		ToSpirit:    toSpirit,
		WireBytes:   wireBytes,
		Broadcast:   broadcast,
		ToOld:       toOld,
		ToOldAndNew: toOldAndNew,
		MagicKeyId:  magicKeyId,
		SummoningId: summoningId,
	}
}

func (msg *MsgCreateMeditation) Route() string {
	return RouterKey
}

func (msg *MsgCreateMeditation) Type() string {
	return TypeMsgCreateMeditation
}

func (msg *MsgCreateMeditation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMeditation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMeditation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMeditation{}

func NewMsgDeleteMeditation(creator string, id uint64) *MsgDeleteMeditation {
	return &MsgDeleteMeditation{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMeditation) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMeditation) Type() string {
	return TypeMsgDeleteMeditation
}

func (msg *MsgDeleteMeditation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMeditation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMeditation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
