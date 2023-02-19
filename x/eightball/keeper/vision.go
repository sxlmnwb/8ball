package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVisionCount get the total number of vision
func (k Keeper) GetVisionCount(ctx sdk.Context, magKeyId uint64) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VisionCountKey + magKeyIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVisionCount set the total number of vision
func (k Keeper) SetVisionCount(ctx sdk.Context, magKeyId uint64, count uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VisionCountKey + magKeyIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVision appends a vision in the store with a new id and update the count
func (k Keeper) AppendVision(
	ctx sdk.Context,
	magKeyId uint64,
	vision types.Vision,
) uint64 {
	// Create the vision
	count := k.GetVisionCount(ctx, magKeyId)

	// Set the ID of the appended value
	vision.Id = count

	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VisionKey+magKeyIdStr))
	appendedValue := k.cdc.MustMarshal(&vision)
	store.Set(GetVisionIDBytes(vision.Id), appendedValue)

	// Update vision count
	k.SetVisionCount(ctx, magKeyId, count+1)

	return count
}

// SetVision set a specific vision in the store
func (k Keeper) SetVision(ctx sdk.Context, magKeyId uint64, vision types.Vision) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VisionKey+magKeyIdStr))
	b := k.cdc.MustMarshal(&vision)
	store.Set(GetVisionIDBytes(vision.Id), b)
}

// GetVision returns a vision from its id
func (k Keeper) GetVision(ctx sdk.Context, magKeyId uint64, id uint64) (val types.Vision, found bool) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VisionKey+magKeyIdStr))
	b := store.Get(GetVisionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVision removes a vision from the store
func (k Keeper) RemoveVision(ctx sdk.Context, magKeyId uint64, id uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VisionKey+magKeyIdStr))
	store.Delete(GetVisionIDBytes(id))
}

// GetAllVision returns all vision
func (k Keeper) GetAllVision(ctx sdk.Context, magKeyId uint64) (list []types.Vision) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.VisionKey+magKeyIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vision
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVisionIDBytes returns the byte representation of the ID
func GetVisionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVisionIDFromBytes returns ID in uint64 format from a byte array
func GetVisionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
