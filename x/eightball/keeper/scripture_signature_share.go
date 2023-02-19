package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetScriptureSignatureShareCount get the total number of scriptureSignatureShare
func (k Keeper) GetScriptureSignatureShareCount(ctx sdk.Context, scriptureIndexString string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScriptureSignatureShareCountKey + scriptureIndexString)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetScriptureSignatureShareCount set the total number of scriptureSignatureShare
func (k Keeper) SetScriptureSignatureShareCount(ctx sdk.Context, scriptureIndexString string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScriptureSignatureShareCountKey + scriptureIndexString)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendScriptureSignatureShare appends a scriptureSignatureShare in the store with a new id and update the count
func (k Keeper) AppendScriptureSignatureShare(
	ctx sdk.Context,
	scriptureIndexString string,
	scriptureSignatureShare types.ScriptureSignatureShare,
) uint64 {
	// Create the scriptureSignatureShare
	count := k.GetScriptureSignatureShareCount(ctx, scriptureIndexString)

	// Set the ID of the appended value
	scriptureSignatureShare.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureShareKey+scriptureIndexString))
	appendedValue := k.cdc.MustMarshal(&scriptureSignatureShare)
	store.Set(GetScriptureSignatureShareIDBytes(scriptureSignatureShare.Id), appendedValue)

	// Update scriptureSignatureShare count
	k.SetScriptureSignatureShareCount(ctx, scriptureIndexString, count+1)

	return count
}

// SetScriptureSignatureShare set a specific scriptureSignatureShare in the store
func (k Keeper) SetScriptureSignatureShare(
	ctx sdk.Context,
	scriptureIndexString string,
	scriptureSignatureShare types.ScriptureSignatureShare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureShareKey+scriptureIndexString))
	b := k.cdc.MustMarshal(&scriptureSignatureShare)
	store.Set(GetScriptureSignatureShareIDBytes(scriptureSignatureShare.Id), b)
}

// GetScriptureSignatureShare returns a scriptureSignatureShare from its id
func (k Keeper) GetScriptureSignatureShare(ctx sdk.Context, scriptureIndexString string, id uint64) (val types.ScriptureSignatureShare, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureShareKey+scriptureIndexString))
	b := store.Get(GetScriptureSignatureShareIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScriptureSignatureShare removes a scriptureSignatureShare from the store
func (k Keeper) RemoveScriptureSignatureShare(ctx sdk.Context, scriptureIndexString string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureShareKey+scriptureIndexString))
	store.Delete(GetScriptureSignatureShareIDBytes(id))
}

// GetAllScriptureSignatureShare returns all scriptureSignatureShare
func (k Keeper) GetAllScriptureSignatureShare(ctx sdk.Context, scriptureIndexString string) (list []types.ScriptureSignatureShare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureShareKey+scriptureIndexString))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ScriptureSignatureShare
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetScriptureSignatureShareIDBytes returns the byte representation of the ID
func GetScriptureSignatureShareIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetScriptureSignatureShareIDFromBytes returns ID in uint64 format from a byte array
func GetScriptureSignatureShareIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
