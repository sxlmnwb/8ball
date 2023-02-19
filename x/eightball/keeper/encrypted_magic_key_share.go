package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEncryptedMagicKeyShare set a specific encryptedMagicKeyShare in the store from its index
func (k Keeper) SetEncryptedMagicKeyShare(ctx sdk.Context, encryptedMagicKeyShare types.EncryptedMagicKeyShare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedMagicKeyShareKeyPrefix))
	b := k.cdc.MustMarshal(&encryptedMagicKeyShare)
	store.Set(types.EncryptedMagicKeyShareKey(
		encryptedMagicKeyShare.Index,
	), b)
}

// GetEncryptedMagicKeyShare returns a encryptedMagicKeyShare from its index
func (k Keeper) GetEncryptedMagicKeyShare(
	ctx sdk.Context,
	index string,

) (val types.EncryptedMagicKeyShare, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedMagicKeyShareKeyPrefix))

	b := store.Get(types.EncryptedMagicKeyShareKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEncryptedMagicKeyShare removes a encryptedMagicKeyShare from the store
func (k Keeper) RemoveEncryptedMagicKeyShare(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedMagicKeyShareKeyPrefix))
	store.Delete(types.EncryptedMagicKeyShareKey(
		index,
	))
}

// GetAllEncryptedMagicKeyShare returns all encryptedMagicKeyShare
func (k Keeper) GetAllEncryptedMagicKeyShare(ctx sdk.Context) (list []types.EncryptedMagicKeyShare) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedMagicKeyShareKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EncryptedMagicKeyShare
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
