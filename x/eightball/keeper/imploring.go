package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetImploringCount get the total number of imploring
func (k Keeper) GetImploringCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ImploringCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetImploringCount set the total number of imploring
func (k Keeper) SetImploringCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ImploringCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendImploring appends a imploring in the store with a new id and update the count
func (k Keeper) AppendImploring(
	ctx sdk.Context,
	imploring types.Imploring,
) uint64 {
	// Create the imploring
	count := k.GetImploringCount(ctx)

	// Set the ID of the appended value
	imploring.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ImploringKey))
	appendedValue := k.cdc.MustMarshal(&imploring)
	store.Set(GetImploringIDBytes(imploring.Id), appendedValue)

	// Update imploring count
	k.SetImploringCount(ctx, count+1)

	return count
}

// SetImploring set a specific imploring in the store
func (k Keeper) SetImploring(ctx sdk.Context, imploring types.Imploring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ImploringKey))
	b := k.cdc.MustMarshal(&imploring)
	store.Set(GetImploringIDBytes(imploring.Id), b)
}

// GetImploring returns a imploring from its id
func (k Keeper) GetImploring(ctx sdk.Context, id uint64) (val types.Imploring, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ImploringKey))
	b := store.Get(GetImploringIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveImploring removes a imploring from the store
func (k Keeper) RemoveImploring(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ImploringKey))
	store.Delete(GetImploringIDBytes(id))
}

// GetAllImploring returns all imploring
func (k Keeper) GetAllImploring(ctx sdk.Context) (list []types.Imploring) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ImploringKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Imploring
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetImploringIDBytes returns the byte representation of the ID
func GetImploringIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetImploringIDFromBytes returns ID in uint64 format from a byte array
func GetImploringIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
