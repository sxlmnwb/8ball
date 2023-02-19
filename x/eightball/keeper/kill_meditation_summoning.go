package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillMeditationSummoningCount get the total number of killMeditationSummoning
func (k Keeper) GetKillMeditationSummoningCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillMeditationSummoningCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillMeditationSummoningCount set the total number of killMeditationSummoning
func (k Keeper) SetKillMeditationSummoningCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillMeditationSummoningCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillMeditationSummoning appends a killMeditationSummoning in the store with a new id and update the count
func (k Keeper) AppendKillMeditationSummoning(
	ctx sdk.Context,
	killMeditationSummoning types.KillMeditationSummoning,
) uint64 {
	// Create the killMeditationSummoning
	count := k.GetKillMeditationSummoningCount(ctx)

	// Set the ID of the appended value
	killMeditationSummoning.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMeditationSummoningKey))
	appendedValue := k.cdc.MustMarshal(&killMeditationSummoning)
	store.Set(GetKillMeditationSummoningIDBytes(killMeditationSummoning.Id), appendedValue)

	// Update killMeditationSummoning count
	k.SetKillMeditationSummoningCount(ctx, count+1)

	return count
}

// SetKillMeditationSummoning set a specific killMeditationSummoning in the store
func (k Keeper) SetKillMeditationSummoning(ctx sdk.Context, killMeditationSummoning types.KillMeditationSummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMeditationSummoningKey))
	b := k.cdc.MustMarshal(&killMeditationSummoning)
	store.Set(GetKillMeditationSummoningIDBytes(killMeditationSummoning.Id), b)
}

// GetKillMeditationSummoning returns a killMeditationSummoning from its id
func (k Keeper) GetKillMeditationSummoning(ctx sdk.Context, id uint64) (val types.KillMeditationSummoning, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMeditationSummoningKey))
	b := store.Get(GetKillMeditationSummoningIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillMeditationSummoning removes a killMeditationSummoning from the store
func (k Keeper) RemoveKillMeditationSummoning(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMeditationSummoningKey))
	store.Delete(GetKillMeditationSummoningIDBytes(id))
}

// GetAllKillMeditationSummoning returns all killMeditationSummoning
func (k Keeper) GetAllKillMeditationSummoning(ctx sdk.Context) (list []types.KillMeditationSummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMeditationSummoningKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillMeditationSummoning
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillMeditationSummoningIDBytes returns the byte representation of the ID
func GetKillMeditationSummoningIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillMeditationSummoningIDFromBytes returns ID in uint64 format from a byte array
func GetKillMeditationSummoningIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
