package types

import (
	"testing"

	"eightball/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateKillMeditationSummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateKillMeditationSummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateKillMeditationSummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateKillMeditationSummoning{
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

func TestMsgUpdateKillMeditationSummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKillMeditationSummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKillMeditationSummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKillMeditationSummoning{
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

func TestMsgDeleteKillMeditationSummoning_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKillMeditationSummoning
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKillMeditationSummoning{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKillMeditationSummoning{
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
