package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSignatureShareCount get the total number of signatureShare
func (k Keeper) GetSignatureShareCount(ctx sdk.Context, messageIdUint uint64) uint64 {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignatureShareCountKey + messageIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSignatureShareCount set the total number of signatureShare
func (k Keeper) SetSignatureShareCount(ctx sdk.Context, messageIdUint uint64, count uint64) {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignatureShareCountKey + messageIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSignatureShare appends a signatureShare in the store with a new id and update the count
func (k Keeper) AppendSignatureShare(
	ctx sdk.Context,
	messageIdUint uint64,
	signatureShare types.SignatureShare,
) uint64 {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	// Create the signatureShare
	count := k.GetSignatureShareCount(ctx, messageIdUint)

	// Set the ID of the appended value
	signatureShare.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureShareKey+messageIdStr))
	appendedValue := k.cdc.MustMarshal(&signatureShare)
	store.Set(GetSignatureShareIDBytes(signatureShare.Id), appendedValue)

	// Update signatureShare count
	k.SetSignatureShareCount(ctx, messageIdUint, count+1)

	return count
}

// SetSignatureShare set a specific signatureShare in the store
func (k Keeper) SetSignatureShare(ctx sdk.Context, messageIdUint uint64, signatureShare types.SignatureShare) {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureShareKey+messageIdStr))
	b := k.cdc.MustMarshal(&signatureShare)
	store.Set(GetSignatureShareIDBytes(signatureShare.Id), b)
}

// GetSignatureShare returns a signatureShare from its id
func (k Keeper) GetSignatureShare(ctx sdk.Context, messageIdUint uint64, id uint64) (val types.SignatureShare, found bool) {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureShareKey+messageIdStr))
	b := store.Get(GetSignatureShareIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignatureShare removes a signatureShare from the store
func (k Keeper) RemoveSignatureShare(ctx sdk.Context, messageIdUint uint64, id uint64) {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureShareKey+messageIdStr))
	store.Delete(GetSignatureShareIDBytes(id))
}

// GetAllSignatureShare returns all signatureShare
func (k Keeper) GetAllSignatureShare(ctx sdk.Context, messageIdUint uint64) (list []types.SignatureShare) {
	messageIdStr := strconv.FormatUint(messageIdUint, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureShareKey+messageIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignatureShare
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSignatureShareIDBytes returns the byte representation of the ID
func GetSignatureShareIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSignatureShareIDFromBytes returns ID in uint64 format from a byte array
func GetSignatureShareIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
