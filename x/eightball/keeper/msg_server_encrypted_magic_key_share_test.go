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

func TestEncryptedMagicKeyShareMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EightballKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateEncryptedMagicKeyShare{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateEncryptedMagicKeyShare(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetEncryptedMagicKeyShare(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestEncryptedMagicKeyShareMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateEncryptedMagicKeyShare
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateEncryptedMagicKeyShare{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EightballKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateEncryptedMagicKeyShare(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateEncryptedMagicKeyShare(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetEncryptedMagicKeyShare(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestEncryptedMagicKeyShareMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteEncryptedMagicKeyShare
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteEncryptedMagicKeyShare{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EightballKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateEncryptedMagicKeyShare(wctx, &types.MsgCreateEncryptedMagicKeyShare{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteEncryptedMagicKeyShare(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetEncryptedMagicKeyShare(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
