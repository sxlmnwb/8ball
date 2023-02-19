package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSignedScriptureList = "create_signed_scripture_list"
	TypeMsgUpdateSignedScriptureList = "update_signed_scripture_list"
	TypeMsgDeleteSignedScriptureList = "delete_signed_scripture_list"
)

var _ sdk.Msg = &MsgCreateSignedScriptureList{}

func NewMsgCreateSignedScriptureList(creator string, scriptureIndex string) *MsgCreateSignedScriptureList {
	return &MsgCreateSignedScriptureList{
		Creator:        creator,
		ScriptureIndex: scriptureIndex,
	}
}

func (msg *MsgCreateSignedScriptureList) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignedScriptureList) Type() string {
	return TypeMsgCreateSignedScriptureList
}

func (msg *MsgCreateSignedScriptureList) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignedScriptureList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignedScriptureList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
