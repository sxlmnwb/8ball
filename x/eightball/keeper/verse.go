package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVerseCount get the total number of verse
func (k Keeper) GetVerseCount(ctx sdk.Context, magKeyId uint64) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VerseCountKey + magKeyIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVerseCount set the total number of verse
func (k Keeper) SetVerseCount(ctx sdk.Context, magKeyId uint64, count uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VerseCountKey + magKeyIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVerse appends a verse in the store with a new id and update the count
func (k Keeper) AppendVerse(
	ctx sdk.Context,
	magKeyId uint64,
	verse types.Verse,
) uint64 {
	// Create the verse
	count := k.GetVerseCount(ctx, magKeyId)

	// Set the ID of the appended value
	verse.Id = count

	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerseKey+magKeyIdStr))
	appendedValue := k.cdc.MustMarshal(&verse)
	store.Set(GetVerseIDBytes(verse.Id), appendedValue)

	// Update verse count
	k.SetVerseCount(ctx, magKeyId, count+1)

	return count
}

// SetVerse set a specific verse in the store
func (k Keeper) SetVerse(ctx sdk.Context, magKeyId uint64, verse types.Verse) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerseKey+magKeyIdStr))
	b := k.cdc.MustMarshal(&verse)
	store.Set(GetVerseIDBytes(verse.Id), b)
}

// GetVerse returns a verse from its id
func (k Keeper) GetVerse(ctx sdk.Context, magKeyId uint64, id uint64) (val types.Verse, found bool) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerseKey+magKeyIdStr))
	b := store.Get(GetVerseIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVerse removes a verse from the store
func (k Keeper) RemoveVerse(ctx sdk.Context, magKeyId uint64, id uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerseKey+magKeyIdStr))
	store.Delete(GetVerseIDBytes(id))
}

// GetAllVerse returns all verse
func (k Keeper) GetAllVerse(ctx sdk.Context, magKeyId uint64) (list []types.Verse) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerseKey+magKeyIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Verse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVerseIDBytes returns the byte representation of the ID
func GetVerseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVerseIDFromBytes returns ID in uint64 format from a byte array
func GetVerseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
