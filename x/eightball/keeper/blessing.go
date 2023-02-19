package keeper

import (
	"encoding/binary"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBlessingCount get the total number of blessing
func (k Keeper) GetBlessingCount(ctx sdk.Context, index string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BlessingCountKey + index)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBlessingCount set the total number of blessing
func (k Keeper) SetBlessingCount(ctx sdk.Context, index string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BlessingCountKey + index)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBlessing appends a blessing in the store with a new id and update the count
func (k Keeper) AppendBlessing(
	ctx sdk.Context,
	index string,
	blessing types.Blessing,
) uint64 {
	// Create the blessing
	count := k.GetBlessingCount(ctx, index)

	// Set the ID of the appended value
	blessing.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlessingKey+index))
	appendedValue := k.cdc.MustMarshal(&blessing)
	store.Set(GetBlessingIDBytes(blessing.Id), appendedValue)

	// Update blessing count
	k.SetBlessingCount(ctx, index, count+1)

	return count
}

// SetBlessing set a specific blessing in the store
func (k Keeper) SetBlessing(ctx sdk.Context, index string, blessing types.Blessing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.BlessingKey+index))
	b := k.cdc.MustMarshal(&blessing)
	store.Set(GetBlessingIDBytes(blessing.Id), b)
}

// GetBlessing returns a blessing from its id
func (k Keeper) GetBlessing(ctx sdk.Context, index string, id uint64) (val types.Blessing, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.BlessingKey+index))
	b := store.Get(GetBlessingIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBlessing removes a blessing from the store
func (k Keeper) RemoveBlessing(ctx sdk.Context, index string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.BlessingKey+index))
	store.Delete(GetBlessingIDBytes(id))
}

// GetAllBlessing returns all blessing
func (k Keeper) GetAllBlessing(ctx sdk.Context, index string) (list []types.Blessing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.BlessingKey+index))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Blessing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBlessingIDBytes returns the byte representation of the ID
func GetBlessingIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBlessingIDFromBytes returns ID in uint64 format from a byte array
func GetBlessingIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
