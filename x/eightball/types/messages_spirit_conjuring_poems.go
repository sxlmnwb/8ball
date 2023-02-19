package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSpiritConjuringPoems = "create_spirit_conjuring_poems"
	TypeMsgUpdateSpiritConjuringPoems = "update_spirit_conjuring_poems"
	TypeMsgDeleteSpiritConjuringPoems = "delete_spirit_conjuring_poems"
)

var _ sdk.Msg = &MsgCreateSpiritConjuringPoems{}

func NewMsgCreateSpiritConjuringPoems(creator string, conjuringId uint64, poem string) *MsgCreateSpiritConjuringPoems {
	return &MsgCreateSpiritConjuringPoems{
		Creator:     creator,
		Poem:        poem,
		ConjuringId: conjuringId,
	}
}

func (msg *MsgCreateSpiritConjuringPoems) Route() string {
	return RouterKey
}

func (msg *MsgCreateSpiritConjuringPoems) Type() string {
	return TypeMsgCreateSpiritConjuringPoems
}

func (msg *MsgCreateSpiritConjuringPoems) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSpiritConjuringPoems) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSpiritConjuringPoems) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
