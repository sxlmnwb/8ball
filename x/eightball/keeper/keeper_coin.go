package keeper

import (
	"eightball/x/eightball/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SendCoins(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error {
	return k.bankKeeper.SendCoins(ctx, from, to, coins)
}

func (k Keeper) AddCoins(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)
}

func (k Keeper) SendFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins []sdk.Coin) error {
	eightballCoins := make(sdk.Coins, len(coins))
	for i, c := range coins {
		value, ok := sdk.NewIntFromString(c.Amount.String())
		if !ok {
			err := fmt.Errorf("invalid amount string: %s", c.Amount.String())
			return err
		}
		eightballCoins[i] = sdk.NewCoin(c.Denom, value)
	}

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, to, eightballCoins)
}

func (k Keeper) SendFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins []sdk.Coin) error {
	eightballCoins := make(sdk.Coins, len(coins))
	for i, c := range coins {
		value, ok := sdk.NewIntFromString(c.Amount.String())
		if !ok {
			err := fmt.Errorf("invalid amount string: %s", c.Amount.String())
			return err
		}
		eightballCoins[i] = sdk.NewCoin(c.Denom, value)
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, from, to, eightballCoins)
}

func (k Keeper) BurnFromModule(ctx sdk.Context, module string, coin sdk.Coin) error {
	value, ok := sdk.NewIntFromString(coin.Amount.String())
	if !ok {
		err := fmt.Errorf("invalid amount string: %s", coin.Amount.String())
		return err
	}
	coinsToBurn := sdk.Coins{sdk.Coin{Denom: coin.Denom, Amount: value}}
	err := k.bankKeeper.BurnCoins(ctx, module, coinsToBurn)
	if err != nil {
		return fmt.Errorf("fail to burn assets: %w", err)
	}

	return nil
}

func (k Keeper) MintToModule(ctx sdk.Context, module string, coin sdk.Coin) error {
	value, ok := sdk.NewIntFromString(coin.Amount.String())
	if !ok {
		err := fmt.Errorf("invalid value string: %s", coin.Amount.String())
		return err
	}

	coinsToMint := sdk.Coins{sdk.Coin{Denom: coin.Denom, Amount: value}}
	err := k.bankKeeper.MintCoins(ctx, module, coinsToMint)
	if err != nil {
		return fmt.Errorf("fail to mint assets: %w", err)
	}

	return nil
}

func (k Keeper) MintAndSendToAccount(ctx sdk.Context, to sdk.AccAddress, coin sdk.Coin) error {
	if err := k.MintToModule(ctx, types.ModuleName, coin); err != nil {
		return err
	}

	return k.SendFromModuleToAccount(ctx, types.ModuleName, to, []sdk.Coin{coin})
}

func (k Keeper) GetModuleAccAddress(module string) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(module)
}

func (k Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, addr)
}

func (k Keeper) HasCoins(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) bool {
	balance := k.bankKeeper.GetAllBalances(ctx, addr)
	return balance.IsAllGTE(coins)
}
