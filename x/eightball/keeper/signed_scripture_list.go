package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSignedScriptureListCount get the total number of signedScriptureList
func (k Keeper) GetSignedScriptureListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignedScriptureListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSignedScriptureListCount set the total number of signedScriptureList
func (k Keeper) SetSignedScriptureListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignedScriptureListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSignedScriptureList appends a signedScriptureList in the store with a new id and update the count
func (k Keeper) AppendSignedScriptureList(
	ctx sdk.Context,
	signedScriptureList types.SignedScriptureList,
) uint64 {
	// Create the signedScriptureList
	count := k.GetSignedScriptureListCount(ctx)

	// Set the ID of the appended value
	signedScriptureList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureListKey))
	appendedValue := k.cdc.MustMarshal(&signedScriptureList)
	store.Set(GetSignedScriptureListIDBytes(signedScriptureList.Id), appendedValue)

	// Update signedScriptureList count
	k.SetSignedScriptureListCount(ctx, count+1)

	return count
}

// SetSignedScriptureList set a specific signedScriptureList in the store
func (k Keeper) SetSignedScriptureList(ctx sdk.Context, signedScriptureList types.SignedScriptureList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureListKey))
	b := k.cdc.MustMarshal(&signedScriptureList)
	store.Set(GetSignedScriptureListIDBytes(signedScriptureList.Id), b)
}

// GetSignedScriptureList returns a signedScriptureList from its id
func (k Keeper) GetSignedScriptureList(ctx sdk.Context, id uint64) (val types.SignedScriptureList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureListKey))
	b := store.Get(GetSignedScriptureListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignedScriptureList removes a signedScriptureList from the store
func (k Keeper) RemoveSignedScriptureList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureListKey))
	store.Delete(GetSignedScriptureListIDBytes(id))
}

// GetAllSignedScriptureList returns all signedScriptureList
func (k Keeper) GetAllSignedScriptureList(ctx sdk.Context) (list []types.SignedScriptureList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignedScriptureList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSignedScriptureListIDBytes returns the byte representation of the ID
func GetSignedScriptureListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSignedScriptureListIDFromBytes returns ID in uint64 format from a byte array
func GetSignedScriptureListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
