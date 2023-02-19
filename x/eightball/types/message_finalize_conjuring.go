package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinalizeConjuring = "finalize_conjuring"

var _ sdk.Msg = &MsgFinalizeConjuring{}

func NewMsgFinalizeConjuring(creator string, conjuringId uint64) *MsgFinalizeConjuring {
	return &MsgFinalizeConjuring{
		Creator:     creator,
		ConjuringId: conjuringId,
	}
}

func (msg *MsgFinalizeConjuring) Route() string {
	return RouterKey
}

func (msg *MsgFinalizeConjuring) Type() string {
	return TypeMsgFinalizeConjuring
}

func (msg *MsgFinalizeConjuring) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalizeConjuring) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalizeConjuring) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
