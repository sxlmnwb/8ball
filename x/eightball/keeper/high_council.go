package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetHighCouncilCount get the total number of highCouncil
func (k Keeper) GetHighCouncilCount(ctx sdk.Context, keyId uint64) uint64 {
	keyIdStr := strconv.FormatUint(keyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HighCouncilCountKey + keyIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHighCouncilCount set the total number of highCouncil
func (k Keeper) SetHighCouncilCount(ctx sdk.Context, keyId uint64, count uint64) {
	keyIdStr := strconv.FormatUint(keyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HighCouncilCountKey + keyIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHighCouncil appends a highCouncil in the store with a new id and update the count
func (k Keeper) AppendHighCouncil(
	ctx sdk.Context,
	keyId uint64,
	highCouncil types.HighCouncil,
) uint64 {
	keyIdStr := strconv.FormatUint(keyId, 10)

	// Create the highCouncil
	count := k.GetHighCouncilCount(ctx, keyId)

	// Set the ID of the appended value
	highCouncil.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilKey+keyIdStr))
	appendedValue := k.cdc.MustMarshal(&highCouncil)
	store.Set(GetHighCouncilIDBytes(highCouncil.Id), appendedValue)

	// Update highCouncil count
	k.SetHighCouncilCount(ctx, keyId, count+1)

	return count
}

// SetHighCouncil set a specific highCouncil in the store
func (k Keeper) SetHighCouncil(ctx sdk.Context, keyId uint64, highCouncil types.HighCouncil) {
	keyIdStr := strconv.FormatUint(keyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilKey+keyIdStr))
	b := k.cdc.MustMarshal(&highCouncil)
	store.Set(GetHighCouncilIDBytes(highCouncil.Id), b)
}

// GetHighCouncil returns a highCouncil from its id
func (k Keeper) GetHighCouncil(ctx sdk.Context, keyId uint64, id uint64) (val types.HighCouncil, found bool) {
	keyIdStr := strconv.FormatUint(keyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilKey+keyIdStr))
	b := store.Get(GetHighCouncilIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHighCouncil removes a highCouncil from the store
func (k Keeper) RemoveHighCouncil(ctx sdk.Context, keyId uint64, id uint64) {
	keyIdStr := strconv.FormatUint(keyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilKey+keyIdStr))
	store.Delete(GetHighCouncilIDBytes(id))
}

// GetAllHighCouncil returns all highCouncil
func (k Keeper) GetAllHighCouncil(ctx sdk.Context, keyId uint64) (list []types.HighCouncil) {
	keyIdStr := strconv.FormatUint(keyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilKey+keyIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HighCouncil
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHighCouncilIDBytes returns the byte representation of the ID
func GetHighCouncilIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHighCouncilIDFromBytes returns ID in uint64 format from a byte array
func GetHighCouncilIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
