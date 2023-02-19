package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillImploringCount get the total number of killImploring
func (k Keeper) GetKillImploringCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillImploringCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillImploringCount set the total number of killImploring
func (k Keeper) SetKillImploringCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillImploringCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillImploring appends a killImploring in the store with a new id and update the count
func (k Keeper) AppendKillImploring(
	ctx sdk.Context,
	killImploring types.KillImploring,
) uint64 {
	// Create the killImploring
	count := k.GetKillImploringCount(ctx)

	// Set the ID of the appended value
	killImploring.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillImploringKey))
	appendedValue := k.cdc.MustMarshal(&killImploring)
	store.Set(GetKillImploringIDBytes(killImploring.Id), appendedValue)

	// Update killImploring count
	k.SetKillImploringCount(ctx, count+1)

	return count
}

// SetKillImploring set a specific killImploring in the store
func (k Keeper) SetKillImploring(ctx sdk.Context, killImploring types.KillImploring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillImploringKey))
	b := k.cdc.MustMarshal(&killImploring)
	store.Set(GetKillImploringIDBytes(killImploring.Id), b)
}

// GetKillImploring returns a killImploring from its id
func (k Keeper) GetKillImploring(ctx sdk.Context, id uint64) (val types.KillImploring, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillImploringKey))
	b := store.Get(GetKillImploringIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillImploring removes a killImploring from the store
func (k Keeper) RemoveKillImploring(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillImploringKey))
	store.Delete(GetKillImploringIDBytes(id))
}

// GetAllKillImploring returns all killImploring
func (k Keeper) GetAllKillImploring(ctx sdk.Context) (list []types.KillImploring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillImploringKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillImploring
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillImploringIDBytes returns the byte representation of the ID
func GetKillImploringIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillImploringIDFromBytes returns ID in uint64 format from a byte array
func GetKillImploringIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
