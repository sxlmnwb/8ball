package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetScripture set a specific scripture in the store from its index
func (k Keeper) SetScripture(ctx sdk.Context, scripture types.Scripture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureKeyPrefix))
	b := k.cdc.MustMarshal(&scripture)
	store.Set(types.ScriptureKey(
		scripture.Index,
	), b)
}

// GetScripture returns a scripture from its index
func (k Keeper) GetScripture(
	ctx sdk.Context,
	index string,

) (val types.Scripture, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureKeyPrefix))

	b := store.Get(types.ScriptureKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScripture removes a scripture from the store
func (k Keeper) RemoveScripture(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureKeyPrefix))
	store.Delete(types.ScriptureKey(
		index,
	))
}

// GetAllScripture returns all scripture
func (k Keeper) GetAllScripture(ctx sdk.Context) (list []types.Scripture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScriptureKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Scripture
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
