package types

import (
	"testing"

	"eightball/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEncryptedMagicKeyShare_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEncryptedMagicKeyShare
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEncryptedMagicKeyShare{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateEncryptedMagicKeyShare{
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

func TestMsgUpdateEncryptedMagicKeyShare_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateEncryptedMagicKeyShare
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateEncryptedMagicKeyShare{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateEncryptedMagicKeyShare{
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

func TestMsgDeleteEncryptedMagicKeyShare_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteEncryptedMagicKeyShare
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteEncryptedMagicKeyShare{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteEncryptedMagicKeyShare{
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
