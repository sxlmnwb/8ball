package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillMagicKeySummoningCount get the total number of killMagicKeySummoning
func (k Keeper) GetKillMagicKeySummoningCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillMagicKeySummoningCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillMagicKeySummoningCount set the total number of killMagicKeySummoning
func (k Keeper) SetKillMagicKeySummoningCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillMagicKeySummoningCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillMagicKeySummoning appends a killMagicKeySummoning in the store with a new id and update the count
func (k Keeper) AppendKillMagicKeySummoning(
	ctx sdk.Context,
	killMagicKeySummoning types.KillMagicKeySummoning,
) uint64 {
	// Create the killMagicKeySummoning
	count := k.GetKillMagicKeySummoningCount(ctx)

	// Set the ID of the appended value
	killMagicKeySummoning.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMagicKeySummoningKey))
	appendedValue := k.cdc.MustMarshal(&killMagicKeySummoning)
	store.Set(GetKillMagicKeySummoningIDBytes(killMagicKeySummoning.Id), appendedValue)

	// Update killMagicKeySummoning count
	k.SetKillMagicKeySummoningCount(ctx, count+1)

	return count
}

// SetKillMagicKeySummoning set a specific killMagicKeySummoning in the store
func (k Keeper) SetKillMagicKeySummoning(ctx sdk.Context, killMagicKeySummoning types.KillMagicKeySummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMagicKeySummoningKey))
	b := k.cdc.MustMarshal(&killMagicKeySummoning)
	store.Set(GetKillMagicKeySummoningIDBytes(killMagicKeySummoning.Id), b)
}

// GetKillMagicKeySummoning returns a killMagicKeySummoning from its id
func (k Keeper) GetKillMagicKeySummoning(ctx sdk.Context, id uint64) (val types.KillMagicKeySummoning, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMagicKeySummoningKey))
	b := store.Get(GetKillMagicKeySummoningIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillMagicKeySummoning removes a killMagicKeySummoning from the store
func (k Keeper) RemoveKillMagicKeySummoning(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMagicKeySummoningKey))
	store.Delete(GetKillMagicKeySummoningIDBytes(id))
}

// GetAllKillMagicKeySummoning returns all killMagicKeySummoning
func (k Keeper) GetAllKillMagicKeySummoning(ctx sdk.Context) (list []types.KillMagicKeySummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillMagicKeySummoningKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillMagicKeySummoning
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillMagicKeySummoningIDBytes returns the byte representation of the ID
func GetKillMagicKeySummoningIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillMagicKeySummoningIDFromBytes returns ID in uint64 format from a byte array
func GetKillMagicKeySummoningIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
