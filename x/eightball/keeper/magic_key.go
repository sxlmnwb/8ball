package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMagicKeyCount get the total number of magicKey
func (k Keeper) GetMagicKeyCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MagicKeyCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMagicKeyCount set the total number of magicKey
func (k Keeper) SetMagicKeyCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MagicKeyCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMagicKey appends a magicKey in the store with a new id and update the count
func (k Keeper) AppendMagicKey(
	ctx sdk.Context,
	magicKey types.MagicKey,
) uint64 {
	// Create the magicKey
	count := k.GetMagicKeyCount(ctx)

	// Set the ID of the appended value
	magicKey.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeyKey))
	appendedValue := k.cdc.MustMarshal(&magicKey)
	store.Set(GetMagicKeyIDBytes(magicKey.Id), appendedValue)

	// Update magicKey count
	k.SetMagicKeyCount(ctx, count+1)

	return count
}

// SetMagicKey set a specific magicKey in the store
func (k Keeper) SetMagicKey(ctx sdk.Context, magicKey types.MagicKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeyKey))
	b := k.cdc.MustMarshal(&magicKey)
	store.Set(GetMagicKeyIDBytes(magicKey.Id), b)
}

// GetMagicKey returns a magicKey from its id
func (k Keeper) GetMagicKey(ctx sdk.Context, id uint64) (val types.MagicKey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeyKey))
	b := store.Get(GetMagicKeyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMagicKey removes a magicKey from the store
func (k Keeper) RemoveMagicKey(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeyKey))
	store.Delete(GetMagicKeyIDBytes(id))
}

// GetAllMagicKey returns all magicKey
func (k Keeper) GetAllMagicKey(ctx sdk.Context) (list []types.MagicKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MagicKeyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MagicKey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMagicKeyIDBytes returns the byte representation of the ID
func GetMagicKeyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMagicKeyIDFromBytes returns ID in uint64 format from a byte array
func GetMagicKeyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
