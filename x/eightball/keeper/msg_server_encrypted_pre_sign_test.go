package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "eightball/testutil/keeper"
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEncryptedPreSignMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EightballKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateEncryptedPreSign{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateEncryptedPreSign(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetEncryptedPreSign(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestEncryptedPreSignMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateEncryptedPreSign
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateEncryptedPreSign{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EightballKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateEncryptedPreSign(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateEncryptedPreSign(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetEncryptedPreSign(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestEncryptedPreSignMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteEncryptedPreSign
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteEncryptedPreSign{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EightballKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateEncryptedPreSign(wctx, &types.MsgCreateEncryptedPreSign{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteEncryptedPreSign(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetEncryptedPreSign(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
