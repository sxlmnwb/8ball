package eightball

import (
	"eightball/x/eightball/keeper"
	"eightball/x/eightball/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set all the killConjuring
	for _, elem := range genState.KillConjuringList {
		k.SetKillConjuring(ctx, elem)
	}

	// Set killConjuring count
	k.SetKillConjuringCount(ctx, genState.KillConjuringCount)
	// Set all the killImploring
	for _, elem := range genState.KillImploringList {
		k.SetKillImploring(ctx, elem)
	}

	// Set killImploring count
	k.SetKillImploringCount(ctx, genState.KillImploringCount)
	// Set all the killMeditationSummoning
	for _, elem := range genState.KillMeditationSummoningList {
		k.SetKillMeditationSummoning(ctx, elem)
	}

	// Set killMeditationSummoning count
	k.SetKillMeditationSummoningCount(ctx, genState.KillMeditationSummoningCount)
	// Set all the killMagicKeySummoning
	for _, elem := range genState.KillMagicKeySummoningList {
		k.SetKillMagicKeySummoning(ctx, elem)
	}

	// Set killMagicKeySummoning count
	k.SetKillMagicKeySummoningCount(ctx, genState.KillMagicKeySummoningCount)
	// Set all the killScriptureSignatureRequest
	for _, elem := range genState.KillScriptureSignatureRequestList {
		k.SetKillScriptureSignatureRequest(ctx, elem)
	}

	// Set killScriptureSignatureRequest count
	k.SetKillScriptureSignatureRequestCount(ctx, genState.KillScriptureSignatureRequestCount)
	// Set all the killSignatureRequest
	for _, elem := range genState.KillSignatureRequestList {
		k.SetKillSignatureRequest(ctx, elem)
	}

	// Set killSignatureRequest count
	k.SetKillSignatureRequestCount(ctx, genState.KillSignatureRequestCount)
	// Set all the encryptedMagicKeyShare
	for _, elem := range genState.EncryptedMagicKeyShareList {
		k.SetEncryptedMagicKeyShare(ctx, elem)
	}
	// Set all the encryptedPreSign
	for _, elem := range genState.EncryptedPreSignList {
		k.SetEncryptedPreSign(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.KillConjuringList = k.GetAllKillConjuring(ctx)
	genesis.KillConjuringCount = k.GetKillConjuringCount(ctx)
	genesis.KillImploringList = k.GetAllKillImploring(ctx)
	genesis.KillImploringCount = k.GetKillImploringCount(ctx)
	genesis.KillMeditationSummoningList = k.GetAllKillMeditationSummoning(ctx)
	genesis.KillMeditationSummoningCount = k.GetKillMeditationSummoningCount(ctx)
	genesis.KillMagicKeySummoningList = k.GetAllKillMagicKeySummoning(ctx)
	genesis.KillMagicKeySummoningCount = k.GetKillMagicKeySummoningCount(ctx)
	genesis.KillScriptureSignatureRequestList = k.GetAllKillScriptureSignatureRequest(ctx)
	genesis.KillScriptureSignatureRequestCount = k.GetKillScriptureSignatureRequestCount(ctx)
	genesis.KillSignatureRequestList = k.GetAllKillSignatureRequest(ctx)
	genesis.KillSignatureRequestCount = k.GetKillSignatureRequestCount(ctx)
	genesis.EncryptedMagicKeyShareList = k.GetAllEncryptedMagicKeyShare(ctx)
	genesis.EncryptedPreSignList = k.GetAllEncryptedPreSign(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
