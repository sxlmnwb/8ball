package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVerse = "create_verse"
	TypeMsgDeleteVerse = "delete_verse"
)

var _ sdk.Msg = &MsgCreateVerse{}

func NewMsgCreateVerse(creator string, fromSpirit string, toSpirit []string, wireBytes string, broadcast bool, toOld bool, toOldAndNew bool, magicKeyId uint64, summoningId uint64) *MsgCreateVerse {
	return &MsgCreateVerse{
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

func (msg *MsgCreateVerse) Route() string {
	return RouterKey
}

func (msg *MsgCreateVerse) Type() string {
	return TypeMsgCreateVerse
}

func (msg *MsgCreateVerse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVerse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVerse{}

func NewMsgDeleteVerse(creator string, id uint64) *MsgDeleteVerse {
	return &MsgDeleteVerse{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteVerse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVerse) Type() string {
	return TypeMsgDeleteVerse
}

func (msg *MsgDeleteVerse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVerse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
