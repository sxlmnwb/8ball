package keeper

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSignedScripture(goCtx context.Context,
	msg *types.MsgCreateSignedScripture) (*types.MsgCreateSignedScriptureResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// require request to come from bonded spirit
	creator, err := sdk.AccAddressFromBech32(msg.Witness)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Witness))
	}
	thisSpiritPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisSpiritPower = thisSpiritPower.Quo(math.NewInt(1000000))
	if thisSpiritPower.IsNil() || thisSpiritPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Witness))
	}

	// Check if the value already exists
	_, isFound := k.GetSignedScripture(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"index already set")
	}

	// check the message signature is correct

	currMagicKey, found := k.GetCurrentMagicKey(ctx)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"can not get current magic key")
	}

	magKey, found := k.GetMagicKey(ctx, currMagicKey.CurrentMagicKey)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"can not get current magic key info")
	}

	sigBytes, err := base64.StdEncoding.DecodeString(msg.Signature)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			"can not parse signature btyes")
	}

	buffer := bytes.Buffer{}
	buffer.Grow(wire.VarIntSerializeSize(uint64(len(msg.Index))))

	// varIntProtoVer is the protocol version to use for serializing
	// N as a VarInt
	// Copied from
	// https://github.com/btcsuite/btcutil/blob/master/gcs/gcs.go#L37
	const varIntProtoVer uint32 = 0
	const magicMessage = "\x18Bitcoin Signed Message:\n"

	wire.WriteVarInt(&buffer, varIntProtoVer, uint64(len(msg.Index)))

	messageHash := chainhash.DoubleHashB([]byte(
		magicMessage + buffer.String() + msg.Index))

	btcecPubKey, _, _ := btcec.RecoverCompact(
		btcec.S256(), sigBytes, messageHash)

	verified := false

	if btcecPubKey != nil &&
		btcecPubKey.ToECDSA().X.String() == magKey.EcPointX &&
		btcecPubKey.ToECDSA().Y.String() == magKey.EcPointY {
		verified = true
	}

	if !verified {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			"Message signature did not verify with current magic key")
	}

	if string(msg.BtcAddress[0]) != "1" {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			"btc address must start with 1")

	}

	if magKey.BitcoinAlias != msg.BtcAddress {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrKeyNotFound,
			"signature btc address does not match magic key")
	}

	var signedScripture = types.SignedScripture{
		Index:      msg.Index,
		Signature:  msg.Signature,
		BtcAddress: msg.BtcAddress,
		MagicKeyId: msg.MagicKeyId,
	}

	scripture, ok := k.GetScripture(ctx, msg.Index)

	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not not find scripture with index %s \n",
				scripture))
	}

	k.RemoveScriptureSignatureRequest(ctx, msg.SigRequestId)

	var imploring = types.Imploring{
		Height: uint64(ctx.BlockHeight()),
		Index:  msg.Index,
	}

	// append imploring for spirits to discover and coordinate a blessing
	k.AppendImploring(ctx, imploring)

	k.SetSignedScripture(
		ctx,
		signedScripture,
	)

	return &types.MsgCreateSignedScriptureResponse{}, nil
}
