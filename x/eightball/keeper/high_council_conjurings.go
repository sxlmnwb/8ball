package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetHighCouncilConjuringsCount get the total number of highCouncilConjurings
func (k Keeper) GetHighCouncilConjuringsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HighCouncilConjuringsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHighCouncilConjuringsCount set the total number of highCouncilConjurings
func (k Keeper) SetHighCouncilConjuringsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HighCouncilConjuringsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHighCouncilConjurings appends a highCouncilConjurings in the store with a new id and update the count
func (k Keeper) AppendHighCouncilConjurings(
	ctx sdk.Context,
	highCouncilConjurings types.HighCouncilConjurings,
) uint64 {
	// Create the highCouncilConjurings
	count := k.GetHighCouncilConjuringsCount(ctx)

	// Set the ID of the appended value
	highCouncilConjurings.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilConjuringsKey))
	appendedValue := k.cdc.MustMarshal(&highCouncilConjurings)
	store.Set(GetHighCouncilConjuringsIDBytes(highCouncilConjurings.Id), appendedValue)

	// Update highCouncilConjurings count
	k.SetHighCouncilConjuringsCount(ctx, count+1)

	return count
}

// SetHighCouncilConjurings set a specific highCouncilConjurings in the store
func (k Keeper) SetHighCouncilConjurings(ctx sdk.Context, highCouncilConjurings types.HighCouncilConjurings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilConjuringsKey))
	b := k.cdc.MustMarshal(&highCouncilConjurings)
	store.Set(GetHighCouncilConjuringsIDBytes(highCouncilConjurings.Id), b)
}

// GetHighCouncilConjurings returns a highCouncilConjurings from its id
func (k Keeper) GetHighCouncilConjurings(ctx sdk.Context, id uint64) (val types.HighCouncilConjurings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilConjuringsKey))
	b := store.Get(GetHighCouncilConjuringsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHighCouncilConjurings removes a highCouncilConjurings from the store
func (k Keeper) RemoveHighCouncilConjurings(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilConjuringsKey))
	store.Delete(GetHighCouncilConjuringsIDBytes(id))
}

// GetAllHighCouncilConjurings returns all highCouncilConjurings
func (k Keeper) GetAllHighCouncilConjurings(ctx sdk.Context) (list []types.HighCouncilConjurings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HighCouncilConjuringsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HighCouncilConjurings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHighCouncilConjuringsIDBytes returns the byte representation of the ID
func GetHighCouncilConjuringsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHighCouncilConjuringsIDFromBytes returns ID in uint64 format from a byte array
func GetHighCouncilConjuringsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
