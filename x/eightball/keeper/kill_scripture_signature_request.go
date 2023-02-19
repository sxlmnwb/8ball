package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillScriptureSignatureRequestCount get the total number of killScriptureSignatureRequest
func (k Keeper) GetKillScriptureSignatureRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillScriptureSignatureRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillScriptureSignatureRequestCount set the total number of killScriptureSignatureRequest
func (k Keeper) SetKillScriptureSignatureRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillScriptureSignatureRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillScriptureSignatureRequest appends a killScriptureSignatureRequest in the store with a new id and update the count
func (k Keeper) AppendKillScriptureSignatureRequest(
	ctx sdk.Context,
	killScriptureSignatureRequest types.KillScriptureSignatureRequest,
) uint64 {
	// Create the killScriptureSignatureRequest
	count := k.GetKillScriptureSignatureRequestCount(ctx)

	// Set the ID of the appended value
	killScriptureSignatureRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillScriptureSignatureRequestKey))
	appendedValue := k.cdc.MustMarshal(&killScriptureSignatureRequest)
	store.Set(GetKillScriptureSignatureRequestIDBytes(killScriptureSignatureRequest.Id), appendedValue)

	// Update killScriptureSignatureRequest count
	k.SetKillScriptureSignatureRequestCount(ctx, count+1)

	return count
}

// SetKillScriptureSignatureRequest set a specific killScriptureSignatureRequest in the store
func (k Keeper) SetKillScriptureSignatureRequest(ctx sdk.Context, killScriptureSignatureRequest types.KillScriptureSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillScriptureSignatureRequestKey))
	b := k.cdc.MustMarshal(&killScriptureSignatureRequest)
	store.Set(GetKillScriptureSignatureRequestIDBytes(killScriptureSignatureRequest.Id), b)
}

// GetKillScriptureSignatureRequest returns a killScriptureSignatureRequest from its id
func (k Keeper) GetKillScriptureSignatureRequest(ctx sdk.Context, id uint64) (val types.KillScriptureSignatureRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillScriptureSignatureRequestKey))
	b := store.Get(GetKillScriptureSignatureRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillScriptureSignatureRequest removes a killScriptureSignatureRequest from the store
func (k Keeper) RemoveKillScriptureSignatureRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillScriptureSignatureRequestKey))
	store.Delete(GetKillScriptureSignatureRequestIDBytes(id))
}

// GetAllKillScriptureSignatureRequest returns all killScriptureSignatureRequest
func (k Keeper) GetAllKillScriptureSignatureRequest(ctx sdk.Context) (list []types.KillScriptureSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillScriptureSignatureRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillScriptureSignatureRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillScriptureSignatureRequestIDBytes returns the byte representation of the ID
func GetKillScriptureSignatureRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillScriptureSignatureRequestIDFromBytes returns ID in uint64 format from a byte array
func GetKillScriptureSignatureRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
