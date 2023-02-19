package keeper

import (
	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBlessingReceipt set a specific blessingReceipt in the store from its index
func (k Keeper) SetBlessingReceipt(ctx sdk.Context, blessingReceipt types.BlessingReceipt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlessingReceiptKeyPrefix))
	b := k.cdc.MustMarshal(&blessingReceipt)
	store.Set(types.BlessingReceiptKey(
		blessingReceipt.Index,
	), b)
}

// GetBlessingReceipt returns a blessingReceipt from its index
func (k Keeper) GetBlessingReceipt(
	ctx sdk.Context,
	index string,

) (val types.BlessingReceipt, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlessingReceiptKeyPrefix))

	b := store.Get(types.BlessingReceiptKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBlessingReceipt removes a blessingReceipt from the store
func (k Keeper) RemoveBlessingReceipt(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlessingReceiptKeyPrefix))
	store.Delete(types.BlessingReceiptKey(
		index,
	))
}

// GetAllBlessingReceipt returns all blessingReceipt
func (k Keeper) GetAllBlessingReceipt(ctx sdk.Context) (list []types.BlessingReceipt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlessingReceiptKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BlessingReceipt
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
