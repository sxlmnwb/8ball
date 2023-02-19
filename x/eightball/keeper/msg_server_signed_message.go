package keeper

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"context"

	"eightball/x/eightball/types"

	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSignedMessage(
	goCtx context.Context,
	msg *types.MsgCreateSignedMessage) (*types.MsgCreateSignedMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// require request to come from bonded spirit
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("can not get address from message creator %s \n",
				msg.Creator))
	}
	thisSpiritPower := k.stakingKeeper.GetDelegatorBonded(ctx, creator)
	thisSpiritPower = thisSpiritPower.Quo(math.NewInt(1000000))
	if thisSpiritPower.IsNil() || thisSpiritPower.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound,
			fmt.Sprintf("creator is not a delegator %s \n",
				msg.Creator))
	}

	var signedMessage = types.SignedMessage{
		Creator:        msg.Creator,
		Body:           msg.Body,
		Signature:      msg.Signature,
		BitcoinAddress: msg.BitcoinAddress,
		MessageId:      msg.MessageId,
		MagicKeyId:     msg.MagicKeyId,
	}

	// do not add duplicate signatures
	signedMessageCount := k.GetSignedMessageCount(ctx, msg.MessageId)

	if signedMessageCount > uint64(0) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "signed message already exists")
	}

	// check the message signature is correct

	currMagicKey, found := k.GetCurrentMagicKey(ctx)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "can not get current magic key")

	}

	magKey, found := k.GetMagicKey(ctx, currMagicKey.CurrentMagicKey)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "can not get current magic key info")

	}

	sigBytes, err := base64.StdEncoding.DecodeString(msg.Signature)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "can not parse signature btyes")
	}

	messageBody := msg.Body

	buffer := bytes.Buffer{}
	buffer.Grow(wire.VarIntSerializeSize(uint64(len(messageBody))))
	// varIntProtoVer is the protocol version to use for serializing N as a VarInt
	// Copied from https://github.com/btcsuite/btcutil/blob/master/gcs/gcs.go#L37
	const varIntProtoVer uint32 = 0
	const magicMessage = "\x18Bitcoin Signed Message:\n"

	// If we cannot write the VarInt, just panic since that should never happen
	if err := wire.WriteVarInt(&buffer, varIntProtoVer, uint64(len(messageBody))); err != nil {
		panic(err)
	}

	messageHash := chainhash.DoubleHashB([]byte(magicMessage + buffer.String() + messageBody))

	btcecPubKey, _, _ := btcec.RecoverCompact(btcec.S256(), sigBytes, messageHash)

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

	id := k.AppendSignedMessage(
		ctx,
		msg.MessageId,
		signedMessage,
	)

	k.RemoveSignatureRequest(ctx, msg.SigRequestId)

	return &types.MsgCreateSignedMessageResponse{
		Id: id,
	}, nil
}
