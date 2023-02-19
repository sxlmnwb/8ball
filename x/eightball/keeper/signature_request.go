package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSignatureRequestCount get the total number of signatureRequest
func (k Keeper) GetSignatureRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignatureRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSignatureRequestCount set the total number of signatureRequest
func (k Keeper) SetSignatureRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignatureRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSignatureRequest appends a signatureRequest in the store with a new id and update the count
func (k Keeper) AppendSignatureRequest(
	ctx sdk.Context,
	signatureRequest types.SignatureRequest,
) uint64 {
	// Create the signatureRequest
	count := k.GetSignatureRequestCount(ctx)

	// Set the ID of the appended value
	signatureRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureRequestKey))
	appendedValue := k.cdc.MustMarshal(&signatureRequest)
	store.Set(GetSignatureRequestIDBytes(signatureRequest.Id), appendedValue)

	// Update signatureRequest count
	k.SetSignatureRequestCount(ctx, count+1)

	return count
}

// SetSignatureRequest set a specific signatureRequest in the store
func (k Keeper) SetSignatureRequest(ctx sdk.Context, signatureRequest types.SignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureRequestKey))
	b := k.cdc.MustMarshal(&signatureRequest)
	store.Set(GetSignatureRequestIDBytes(signatureRequest.Id), b)
}

// GetSignatureRequest returns a signatureRequest from its id
func (k Keeper) GetSignatureRequest(ctx sdk.Context, id uint64) (val types.SignatureRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureRequestKey))
	b := store.Get(GetSignatureRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignatureRequest removes a signatureRequest from the store
func (k Keeper) RemoveSignatureRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureRequestKey))
	store.Delete(GetSignatureRequestIDBytes(id))
}

// GetAllSignatureRequest returns all signatureRequest
func (k Keeper) GetAllSignatureRequest(ctx sdk.Context) (list []types.SignatureRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignatureRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSignatureRequestIDBytes returns the byte representation of the ID
func GetSignatureRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSignatureRequestIDFromBytes returns ID in uint64 format from a byte array
func GetSignatureRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
