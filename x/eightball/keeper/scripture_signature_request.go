package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetScriptureSignatureRequestCount get the total number of scriptureSignatureRequest
func (k Keeper) GetScriptureSignatureRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScriptureSignatureRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetScriptureSignatureRequestCount set the total number of scriptureSignatureRequest
func (k Keeper) SetScriptureSignatureRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScriptureSignatureRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendScriptureSignatureRequest appends a scriptureSignatureRequest in the store with a new id and update the count
func (k Keeper) AppendScriptureSignatureRequest(
	ctx sdk.Context,
	scriptureSignatureRequest types.ScriptureSignatureRequest,
) uint64 {
	// Create the scriptureSignatureRequest
	count := k.GetScriptureSignatureRequestCount(ctx)

	// Set the ID of the appended value
	scriptureSignatureRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureRequestKey))
	appendedValue := k.cdc.MustMarshal(&scriptureSignatureRequest)
	store.Set(GetScriptureSignatureRequestIDBytes(scriptureSignatureRequest.Id), appendedValue)

	// Update scriptureSignatureRequest count
	k.SetScriptureSignatureRequestCount(ctx, count+1)

	return count
}

// SetScriptureSignatureRequest set a specific scriptureSignatureRequest in the store
func (k Keeper) SetScriptureSignatureRequest(ctx sdk.Context, scriptureSignatureRequest types.ScriptureSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureRequestKey))
	b := k.cdc.MustMarshal(&scriptureSignatureRequest)
	store.Set(GetScriptureSignatureRequestIDBytes(scriptureSignatureRequest.Id), b)
}

// GetScriptureSignatureRequest returns a scriptureSignatureRequest from its id
func (k Keeper) GetScriptureSignatureRequest(ctx sdk.Context, id uint64) (val types.ScriptureSignatureRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureRequestKey))
	b := store.Get(GetScriptureSignatureRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScriptureSignatureRequest removes a scriptureSignatureRequest from the store
func (k Keeper) RemoveScriptureSignatureRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureRequestKey))
	store.Delete(GetScriptureSignatureRequestIDBytes(id))
}

// GetAllScriptureSignatureRequest returns all scriptureSignatureRequest
func (k Keeper) GetAllScriptureSignatureRequest(ctx sdk.Context) (list []types.ScriptureSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureSignatureRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ScriptureSignatureRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetScriptureSignatureRequestIDBytes returns the byte representation of the ID
func GetScriptureSignatureRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetScriptureSignatureRequestIDFromBytes returns ID in uint64 format from a byte array
func GetScriptureSignatureRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
