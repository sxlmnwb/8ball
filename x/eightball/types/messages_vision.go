package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVision = "create_vision"
	TypeMsgUpdateVision = "update_vision"
	TypeMsgDeleteVision = "delete_vision"
)

var _ sdk.Msg = &MsgCreateVision{}

func NewMsgCreateVision(creator string, magicKeyId uint64, ecPointX string, ecPointY string, summoningId uint64) *MsgCreateVision {
	return &MsgCreateVision{
		Creator:     creator,
		MagicKeyId:  magicKeyId,
		EcPointX:    ecPointX,
		EcPointY:    ecPointY,
		SummoningId: summoningId,
	}
}

func (msg *MsgCreateVision) Route() string {
	return RouterKey
}

func (msg *MsgCreateVision) Type() string {
	return TypeMsgCreateVision
}

func (msg *MsgCreateVision) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVision) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVision) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVision{}

func NewMsgDeleteVision(creator string, magicKeyId uint64, id uint64) *MsgDeleteVision {
	return &MsgDeleteVision{
		Id:         id,
		MagicKeyId: magicKeyId,
		Creator:    creator,
	}
}
func (msg *MsgDeleteVision) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVision) Type() string {
	return TypeMsgDeleteVision
}

func (msg *MsgDeleteVision) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVision) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVision) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
