package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEncryptedPreSign set a specific encryptedPreSign in the store from its index
func (k Keeper) SetEncryptedPreSign(ctx sdk.Context, encryptedPreSign types.EncryptedPreSign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedPreSignKeyPrefix))
	b := k.cdc.MustMarshal(&encryptedPreSign)
	store.Set(types.EncryptedPreSignKey(
		encryptedPreSign.Index,
	), b)
}

// GetEncryptedPreSign returns a encryptedPreSign from its index
func (k Keeper) GetEncryptedPreSign(
	ctx sdk.Context,
	index string,

) (val types.EncryptedPreSign, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedPreSignKeyPrefix))

	b := store.Get(types.EncryptedPreSignKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEncryptedPreSign removes a encryptedPreSign from the store
func (k Keeper) RemoveEncryptedPreSign(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedPreSignKeyPrefix))
	store.Delete(types.EncryptedPreSignKey(
		index,
	))
}

// GetAllEncryptedPreSign returns all encryptedPreSign
func (k Keeper) GetAllEncryptedPreSign(ctx sdk.Context) (list []types.EncryptedPreSign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedPreSignKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EncryptedPreSign
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
