package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCurrentMagicKey set currentMagicKey in the store
func (k Keeper) SetCurrentMagicKey(ctx sdk.Context, currentMagicKey types.CurrentMagicKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrentMagicKeyKey))
	b := k.cdc.MustMarshal(&currentMagicKey)
	store.Set([]byte{0}, b)
}

// GetCurrentMagicKey returns currentMagicKey
func (k Keeper) GetCurrentMagicKey(ctx sdk.Context) (val types.CurrentMagicKey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrentMagicKeyKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCurrentMagicKey removes currentMagicKey from the store
func (k Keeper) RemoveCurrentMagicKey(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrentMagicKeyKey))
	store.Delete([]byte{0})
}
