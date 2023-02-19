package keeper

import (
	"encoding/binary"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMeditationCount get the total number of meditation
func (k Keeper) GetMeditationCount(ctx sdk.Context, magKeyId uint64) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MeditationCountKey + magKeyIdStr)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMeditationCount set the total number of meditation
func (k Keeper) SetMeditationCount(ctx sdk.Context, magKeyId uint64, count uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MeditationCountKey + magKeyIdStr)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMeditation appends a meditation in the store with a new id and update the count
func (k Keeper) AppendMeditation(
	ctx sdk.Context,
	magKeyId uint64,
	meditation types.Meditation,
) uint64 {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	// Create the meditation
	count := k.GetMeditationCount(ctx, magKeyId)

	// Set the ID of the appended value
	meditation.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationKey+magKeyIdStr))
	appendedValue := k.cdc.MustMarshal(&meditation)
	store.Set(GetMeditationIDBytes(meditation.Id), appendedValue)

	// Update meditation count
	k.SetMeditationCount(ctx, magKeyId, count+1)

	return count
}

// SetMeditation set a specific meditation in the store
func (k Keeper) SetMeditation(ctx sdk.Context, magKeyId uint64, meditation types.Meditation) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationKey+magKeyIdStr))
	b := k.cdc.MustMarshal(&meditation)
	store.Set(GetMeditationIDBytes(meditation.Id), b)
}

// GetMeditation returns a meditation from its id
func (k Keeper) GetMeditation(ctx sdk.Context, magKeyId uint64, id uint64) (val types.Meditation, found bool) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationKey+magKeyIdStr))
	b := store.Get(GetMeditationIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMeditation removes a meditation from the store
func (k Keeper) RemoveMeditation(ctx sdk.Context, magKeyId uint64, id uint64) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationKey+magKeyIdStr))
	store.Delete(GetMeditationIDBytes(id))
}

// GetAllMeditation returns all meditation
func (k Keeper) GetAllMeditation(ctx sdk.Context, magKeyId uint64) (list []types.Meditation) {
	magKeyIdStr := strconv.FormatUint(magKeyId, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MeditationKey+magKeyIdStr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Meditation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMeditationIDBytes returns the byte representation of the ID
func GetMeditationIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMeditationIDFromBytes returns ID in uint64 format from a byte array
func GetMeditationIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
