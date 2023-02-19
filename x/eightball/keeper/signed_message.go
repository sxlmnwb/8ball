package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSignedMessageCount get the total number of signedMessage
func (k Keeper) GetSignedMessageCount(ctx sdk.Context, msgId uint64) uint64 {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignedMessageCountKey + msgIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSignedMessageCount set the total number of signedMessage
func (k Keeper) SetSignedMessageCount(ctx sdk.Context, msgId uint64, count uint64) {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignedMessageCountKey + msgIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSignedMessage appends a signedMessage in the store with a new id and update the count
func (k Keeper) AppendSignedMessage(
	ctx sdk.Context,
	msgId uint64,
	signedMessage types.SignedMessage,
) uint64 {
	msgIdStr := strconv.FormatUint(msgId, 10)

	// Create the signedMessage
	count := k.GetSignedMessageCount(ctx, msgId)

	// Set the ID of the appended value
	signedMessage.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedMessageKey+msgIdStr))
	appendedValue := k.cdc.MustMarshal(&signedMessage)
	store.Set(GetSignedMessageIDBytes(signedMessage.Id), appendedValue)

	// Update signedMessage count
	k.SetSignedMessageCount(ctx, msgId, count+1)

	return count
}

// SetSignedMessage set a specific signedMessage in the store
func (k Keeper) SetSignedMessage(ctx sdk.Context, msgId uint64, signedMessage types.SignedMessage) {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedMessageKey+msgIdStr))
	b := k.cdc.MustMarshal(&signedMessage)
	store.Set(GetSignedMessageIDBytes(signedMessage.Id), b)
}

// GetSignedMessage returns a signedMessage from its id
func (k Keeper) GetSignedMessage(ctx sdk.Context, msgId uint64, id uint64) (val types.SignedMessage, found bool) {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedMessageKey+msgIdStr))
	b := store.Get(GetSignedMessageIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignedMessage removes a signedMessage from the store
func (k Keeper) RemoveSignedMessage(ctx sdk.Context, msgId uint64, id uint64) {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedMessageKey+msgIdStr))
	store.Delete(GetSignedMessageIDBytes(id))
}

// GetAllSignedMessage returns all signedMessage
func (k Keeper) GetAllSignedMessage(ctx sdk.Context, msgId uint64) (list []types.SignedMessage) {
	msgIdStr := strconv.FormatUint(msgId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedMessageKey+msgIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignedMessage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSignedMessageIDBytes returns the byte representation of the ID
func GetSignedMessageIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSignedMessageIDFromBytes returns ID in uint64 format from a byte array
func GetSignedMessageIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
