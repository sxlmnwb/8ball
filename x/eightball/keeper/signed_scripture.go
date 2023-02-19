package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSignedScripture set a specific signedScripture in the store from its index
func (k Keeper) SetSignedScripture(ctx sdk.Context, signedScripture types.SignedScripture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureKeyPrefix))
	b := k.cdc.MustMarshal(&signedScripture)
	store.Set(types.SignedScriptureKey(
		signedScripture.Index,
	), b)
}

// GetSignedScripture returns a signedScripture from its index
func (k Keeper) GetSignedScripture(
	ctx sdk.Context,
	index string,

) (val types.SignedScripture, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureKeyPrefix))

	b := store.Get(types.SignedScriptureKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignedScripture removes a signedScripture from the store
func (k Keeper) RemoveSignedScripture(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureKeyPrefix))
	store.Delete(types.SignedScriptureKey(
		index,
	))
}

// GetAllSignedScripture returns all signedScripture
func (k Keeper) GetAllSignedScripture(ctx sdk.Context) (list []types.SignedScripture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignedScriptureKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignedScripture
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
