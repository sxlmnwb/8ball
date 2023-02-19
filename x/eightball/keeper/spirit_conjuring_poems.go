package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSpiritConjuringPoemsCount get the total number of spiritConjuringPoems
func (k Keeper) GetSpiritConjuringPoemsCount(ctx sdk.Context, magKeyId uint64) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SpiritConjuringPoemsCountKey + magKeyIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSpiritConjuringPoemsCount set the total number of spiritConjuringPoems
func (k Keeper) SetSpiritConjuringPoemsCount(ctx sdk.Context, magKeyId uint64, count uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SpiritConjuringPoemsCountKey + magKeyIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSpiritConjuringPoems appends a spiritConjuringPoems in the store with a new id and update the count
func (k Keeper) AppendSpiritConjuringPoems(
	ctx sdk.Context,
	magKeyId uint64,
	spiritConjuringPoems types.SpiritConjuringPoems,
) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	// Create the spiritConjuringPoems
	count := k.GetSpiritConjuringPoemsCount(ctx, magKeyId)

	// Set the ID of the appended value
	spiritConjuringPoems.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpiritConjuringPoemsKey+magKeyIdStr))
	appendedValue := k.cdc.MustMarshal(&spiritConjuringPoems)
	store.Set(GetSpiritConjuringPoemsIDBytes(spiritConjuringPoems.Id), appendedValue)

	// Update spiritConjuringPoems count
	k.SetSpiritConjuringPoemsCount(ctx, magKeyId, count+1)

	return count
}

// SetSpiritConjuringPoems set a specific spiritConjuringPoems in the store
func (k Keeper) SetSpiritConjuringPoems(
	ctx sdk.Context,
	magKeyId uint64,
	spiritConjuringPoems types.SpiritConjuringPoems) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpiritConjuringPoemsKey+magKeyIdStr))
	b := k.cdc.MustMarshal(&spiritConjuringPoems)
	store.Set(GetSpiritConjuringPoemsIDBytes(spiritConjuringPoems.Id), b)
}

// GetSpiritConjuringPoems returns a spiritConjuringPoems from its id
func (k Keeper) GetSpiritConjuringPoems(
	ctx sdk.Context,
	magKeyId uint64,
	id uint64) (val types.SpiritConjuringPoems, found bool) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpiritConjuringPoemsKey+magKeyIdStr))
	b := store.Get(GetSpiritConjuringPoemsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSpiritConjuringPoems removes a spiritConjuringPoems from the store
func (k Keeper) RemoveSpiritConjuringPoems(
	ctx sdk.Context,
	magKeyId uint64,
	id uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpiritConjuringPoemsKey+magKeyIdStr))
	store.Delete(GetSpiritConjuringPoemsIDBytes(id))
}

// GetAllSpiritConjuringPoems returns all spiritConjuringPoems
func (k Keeper) GetAllSpiritConjuringPoems(
	ctx sdk.Context,
	magKeyId uint64,
) (list []types.SpiritConjuringPoems) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpiritConjuringPoemsKey+magKeyIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SpiritConjuringPoems
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSpiritConjuringPoemsIDBytes returns the byte representation of the ID
func GetSpiritConjuringPoemsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSpiritConjuringPoemsIDFromBytes returns ID in uint64 format from a byte array
func GetSpiritConjuringPoemsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
