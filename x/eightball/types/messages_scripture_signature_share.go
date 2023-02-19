package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateScriptureSignatureShare = "create_scripture_signature_share"
	TypeMsgUpdateScriptureSignatureShare = "update_scripture_signature_share"
	TypeMsgDeleteScriptureSignatureShare = "delete_scripture_signature_share"
)

var _ sdk.Msg = &MsgCreateScriptureSignatureShare{}

func NewMsgCreateScriptureSignatureShare(creator string, scriptureIndex string, shareData string, magicKeyId uint64, pubKey string) *MsgCreateScriptureSignatureShare {
	return &MsgCreateScriptureSignatureShare{
		Creator:        creator,
		ScriptureIndex: scriptureIndex,
		ShareData:      shareData,
		MagicKeyId:     magicKeyId,
		PubKey:         pubKey,
	}
}

func (msg *MsgCreateScriptureSignatureShare) Route() string {
	return RouterKey
}

func (msg *MsgCreateScriptureSignatureShare) Type() string {
	return TypeMsgCreateScriptureSignatureShare
}

func (msg *MsgCreateScriptureSignatureShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateScriptureSignatureShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateScriptureSignatureShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
