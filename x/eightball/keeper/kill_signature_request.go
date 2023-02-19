package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKillSignatureRequestCount get the total number of killSignatureRequest
func (k Keeper) GetKillSignatureRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillSignatureRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKillSignatureRequestCount set the total number of killSignatureRequest
func (k Keeper) SetKillSignatureRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KillSignatureRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKillSignatureRequest appends a killSignatureRequest in the store with a new id and update the count
func (k Keeper) AppendKillSignatureRequest(
	ctx sdk.Context,
	killSignatureRequest types.KillSignatureRequest,
) uint64 {
	// Create the killSignatureRequest
	count := k.GetKillSignatureRequestCount(ctx)

	// Set the ID of the appended value
	killSignatureRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillSignatureRequestKey))
	appendedValue := k.cdc.MustMarshal(&killSignatureRequest)
	store.Set(GetKillSignatureRequestIDBytes(killSignatureRequest.Id), appendedValue)

	// Update killSignatureRequest count
	k.SetKillSignatureRequestCount(ctx, count+1)

	return count
}

// SetKillSignatureRequest set a specific killSignatureRequest in the store
func (k Keeper) SetKillSignatureRequest(ctx sdk.Context, killSignatureRequest types.KillSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillSignatureRequestKey))
	b := k.cdc.MustMarshal(&killSignatureRequest)
	store.Set(GetKillSignatureRequestIDBytes(killSignatureRequest.Id), b)
}

// GetKillSignatureRequest returns a killSignatureRequest from its id
func (k Keeper) GetKillSignatureRequest(ctx sdk.Context, id uint64) (val types.KillSignatureRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillSignatureRequestKey))
	b := store.Get(GetKillSignatureRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKillSignatureRequest removes a killSignatureRequest from the store
func (k Keeper) RemoveKillSignatureRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillSignatureRequestKey))
	store.Delete(GetKillSignatureRequestIDBytes(id))
}

// GetAllKillSignatureRequest returns all killSignatureRequest
func (k Keeper) GetAllKillSignatureRequest(ctx sdk.Context) (list []types.KillSignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KillSignatureRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KillSignatureRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKillSignatureRequestIDBytes returns the byte representation of the ID
func GetKillSignatureRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetKillSignatureRequestIDFromBytes returns ID in uint64 format from a byte array
func GetKillSignatureRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
