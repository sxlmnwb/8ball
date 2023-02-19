package types

import (
	"testing"

	"eightball/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateKillMagicKeySummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateKillMagicKeySummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateKillMagicKeySummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateKillMagicKeySummoning{
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

func TestMsgUpdateKillMagicKeySummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKillMagicKeySummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKillMagicKeySummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKillMagicKeySummoning{
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

func TestMsgDeleteKillMagicKeySummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKillMagicKeySummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKillMagicKeySummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKillMagicKeySummoning{
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
