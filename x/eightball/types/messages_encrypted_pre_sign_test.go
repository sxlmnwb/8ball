package types

import (
	"testing"

	"eightball/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEncryptedPreSign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEncryptedPreSign
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEncryptedPreSign{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateEncryptedPreSign{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateEncryptedPreSign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateEncryptedPreSign
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateEncryptedPreSign{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateEncryptedPreSign{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteEncryptedPreSign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteEncryptedPreSign
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteEncryptedPreSign{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteEncryptedPreSign{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
