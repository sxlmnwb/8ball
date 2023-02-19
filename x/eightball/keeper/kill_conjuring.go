package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillConjuringCount get the total number of killConjuring
func (k Keeper) GetKillConjuringCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillConjuringCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillConjuringCount set the total number of killConjuring
func (k Keeper) SetKillConjuringCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillConjuringCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillConjuring appends a killConjuring in the store with a new id and update the count
func (k Keeper) AppendKillConjuring(
	ctx sdk.Context,
	killConjuring types.KillConjuring,
) uint64 {
	// Create the killConjuring
	count := k.GetKillConjuringCount(ctx)

	// Set the ID of the appended value
	killConjuring.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillConjuringKey))
	appendedValue := k.cdc.MustMarshal(&killConjuring)
	store.Set(GetKillConjuringIDBytes(killConjuring.Id), appendedValue)

	// Update killConjuring count
	k.SetKillConjuringCount(ctx, count+1)

	return count
}

// SetKillConjuring set a specific killConjuring in the store
func (k Keeper) SetKillConjuring(ctx sdk.Context, killConjuring types.KillConjuring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillConjuringKey))
	b := k.cdc.MustMarshal(&killConjuring)
	store.Set(GetKillConjuringIDBytes(killConjuring.Id), b)
}

// GetKillConjuring returns a killConjuring from its id
func (k Keeper) GetKillConjuring(ctx sdk.Context, id uint64) (val types.KillConjuring, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillConjuringKey))
	b := store.Get(GetKillConjuringIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillConjuring removes a killConjuring from the store
func (k Keeper) RemoveKillConjuring(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillConjuringKey))
	store.Delete(GetKillConjuringIDBytes(id))
}

// GetAllKillConjuring returns all killConjuring
func (k Keeper) GetAllKillConjuring(ctx sdk.Context) (list []types.KillConjuring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillConjuringKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillConjuring
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillConjuringIDBytes returns the byte representation of the ID
func GetKillConjuringIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillConjuringIDFromBytes returns ID in uint64 format from a byte array
func GetKillConjuringIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
