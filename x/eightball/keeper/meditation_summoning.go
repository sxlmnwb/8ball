package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMeditationSummoningCount get the total number of meditationSummoning
func (k Keeper) GetMeditationSummoningCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MeditationSummoningCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMeditationSummoningCount set the total number of meditationSummoning
func (k Keeper) SetMeditationSummoningCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MeditationSummoningCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMeditationSummoning appends a meditationSummoning in the store with a new id and update the count
func (k Keeper) AppendMeditationSummoning(
	ctx sdk.Context,
	meditationSummoning types.MeditationSummoning,
) uint64 {
	// Create the meditationSummoning
	count := k.GetMeditationSummoningCount(ctx)

	// Set the ID of the appended value
	meditationSummoning.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationSummoningKey))
	appendedValue := k.cdc.MustMarshal(&meditationSummoning)
	store.Set(GetMeditationSummoningIDBytes(meditationSummoning.Id), appendedValue)

	// Update meditationSummoning count
	k.SetMeditationSummoningCount(ctx, count+1)

	return count
}

// SetMeditationSummoning set a specific meditationSummoning in the store
func (k Keeper) SetMeditationSummoning(ctx sdk.Context, meditationSummoning types.MeditationSummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationSummoningKey))
	b := k.cdc.MustMarshal(&meditationSummoning)
	store.Set(GetMeditationSummoningIDBytes(meditationSummoning.Id), b)
}

// GetMeditationSummoning returns a meditationSummoning from its id
func (k Keeper) GetMeditationSummoning(ctx sdk.Context, id uint64) (val types.MeditationSummoning, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationSummoningKey))
	b := store.Get(GetMeditationSummoningIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMeditationSummoning removes a meditationSummoning from the store
func (k Keeper) RemoveMeditationSummoning(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationSummoningKey))
	store.Delete(GetMeditationSummoningIDBytes(id))
}

// GetAllMeditationSummoning returns all meditationSummoning
func (k Keeper) GetAllMeditationSummoning(ctx sdk.Context) (list []types.MeditationSummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationSummoningKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MeditationSummoning
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMeditationSummoningIDBytes returns the byte representation of the ID
func GetMeditationSummoningIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMeditationSummoningIDFromBytes returns ID in uint64 format from a byte array
func GetMeditationSummoningIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
