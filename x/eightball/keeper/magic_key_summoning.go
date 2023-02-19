package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMagicKeySummoningCount get the total number of magicKeySummoning
func (k Keeper) GetMagicKeySummoningCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MagicKeySummoningCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMagicKeySummoningCount set the total number of magicKeySummoning
func (k Keeper) SetMagicKeySummoningCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MagicKeySummoningCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMagicKeySummoning appends a magicKeySummoning in the store with a new id and update the count
func (k Keeper) AppendMagicKeySummoning(
	ctx sdk.Context,
	magicKeySummoning types.MagicKeySummoning,
) uint64 {
	// Create the magicKeySummoning
	count := k.GetMagicKeySummoningCount(ctx)

	// Set the ID of the appended value
	magicKeySummoning.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeySummoningKey))
	appendedValue := k.cdc.MustMarshal(&magicKeySummoning)
	store.Set(GetMagicKeySummoningIDBytes(magicKeySummoning.Id), appendedValue)

	// Update magicKeySummoning count
	k.SetMagicKeySummoningCount(ctx, count+1)

	return count
}

// SetMagicKeySummoning set a specific magicKeySummoning in the store
func (k Keeper) SetMagicKeySummoning(ctx sdk.Context, magicKeySummoning types.MagicKeySummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeySummoningKey))
	b := k.cdc.MustMarshal(&magicKeySummoning)
	store.Set(GetMagicKeySummoningIDBytes(magicKeySummoning.Id), b)
}

// GetMagicKeySummoning returns a magicKeySummoning from its id
func (k Keeper) GetMagicKeySummoning(ctx sdk.Context, id uint64) (val types.MagicKeySummoning, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeySummoningKey))
	b := store.Get(GetMagicKeySummoningIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMagicKeySummoning removes a magicKeySummoning from the store
func (k Keeper) RemoveMagicKeySummoning(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeySummoningKey))
	store.Delete(GetMagicKeySummoningIDBytes(id))
}

// GetAllMagicKeySummoning returns all magicKeySummoning
func (k Keeper) GetAllMagicKeySummoning(ctx sdk.Context) (list []types.MagicKeySummoning) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeySummoningKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MagicKeySummoning
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMagicKeySummoningIDBytes returns the byte representation of the ID
func GetMagicKeySummoningIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMagicKeySummoningIDFromBytes returns ID in uint64 format from a byte array
func GetMagicKeySummoningIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
