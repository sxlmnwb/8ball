package eightball_test

import (
	"testing"

	keepertest "eightball/testutil/keeper"
	"eightball/testutil/nullify"
	"eightball/x/eightball"
	"eightball/x/eightball/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MagicKeyList: []types.MagicKey{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MagicKeyCount: 2,
		MagicKeySummoningList: []types.MagicKeySummoning{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MagicKeySummoningCount: 2,
		HighCouncilList: []types.HighCouncil{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HighCouncilCount: 2,
		HighCouncilConjuringsList: []types.HighCouncilConjurings{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HighCouncilConjuringsCount: 2,
		SpiritConjuringPoemsList: []types.SpiritConjuringPoems{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SpiritConjuringPoemsCount: 2,
		VerseList: []types.Verse{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VerseCount: 2,
		VisionList: []types.Vision{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VisionCount: 2,
		MessageList: []types.Message{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MessageCount: 2,
		SignatureRequestList: []types.SignatureRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SignatureRequestCount: 2,
		SignatureShareList: []types.SignatureShare{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SignatureShareCount: 2,
		SignedMessageList: []types.SignedMessage{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SignedMessageCount: 2,
		MeditationSummoningList: []types.MeditationSummoning{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MeditationSummoningCount: 2,
		MeditationList: []types.Meditation{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MeditationCount: 2,
		ScriptureList: []types.Scripture{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ScriptureSignatureRequestList: []types.ScriptureSignatureRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ScriptureSignatureRequestCount: 2,
		ScriptureSignatureShareList: []types.ScriptureSignatureShare{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ScriptureSignatureShareCount: 2,
		SignedScriptureList: []types.SignedScripture{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SignedScriptureListList: []types.SignedScriptureList{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SignedScriptureListCount: 2,
		BlessingList: []types.Blessing{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BlessingCount: 2,
		BlessingReceiptList: []types.BlessingReceipt{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ImploringList: []types.Imploring{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ImploringCount: 2,
		KillConjuringList: []types.KillConjuring{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillConjuringCount: 2,
		KillImploringList: []types.KillImploring{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillImploringCount: 2,
		KillMeditationSummoningList: []types.KillMeditationSummoning{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillMeditationSummoningCount: 2,
		KillMagicKeySummoningList: []types.KillMagicKeySummoning{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillMagicKeySummoningCount: 2,
		KillScriptureSignatureRequestList: []types.KillScriptureSignatureRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillScriptureSignatureRequestCount: 2,
		KillSignatureRequestList: []types.KillSignatureRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KillSignatureRequestCount: 2,
		EncryptedMagicKeyShareList: []types.EncryptedMagicKeyShare{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		EncryptedPreSignList: []types.EncryptedPreSign{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EightballKeeper(t)
	eightball.InitGenesis(ctx, *k, genesisState)
	got := eightball.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MagicKeyList, got.MagicKeyList)
	require.Equal(t, genesisState.MagicKeyCount, got.MagicKeyCount)
	require.ElementsMatch(t, genesisState.MagicKeySummoningList, got.MagicKeySummoningList)
	require.Equal(t, genesisState.MagicKeySummoningCount, got.MagicKeySummoningCount)
	require.ElementsMatch(t, genesisState.HighCouncilList, got.HighCouncilList)
	require.Equal(t, genesisState.HighCouncilCount, got.HighCouncilCount)
	require.Equal(t, genesisState.CurrentMagicKey, got.CurrentMagicKey)
	require.ElementsMatch(t, genesisState.HighCouncilConjuringsList, got.HighCouncilConjuringsList)
	require.Equal(t, genesisState.HighCouncilConjuringsCount, got.HighCouncilConjuringsCount)
	require.ElementsMatch(t, genesisState.SpiritConjuringPoemsList, got.SpiritConjuringPoemsList)
	require.Equal(t, genesisState.SpiritConjuringPoemsCount, got.SpiritConjuringPoemsCount)
	require.ElementsMatch(t, genesisState.VerseList, got.VerseList)
	require.Equal(t, genesisState.VerseCount, got.VerseCount)
	require.ElementsMatch(t, genesisState.VisionList, got.VisionList)
	require.Equal(t, genesisState.VisionCount, got.VisionCount)
	require.ElementsMatch(t, genesisState.MessageList, got.MessageList)
	require.Equal(t, genesisState.MessageCount, got.MessageCount)
	require.ElementsMatch(t, genesisState.SignatureRequestList, got.SignatureRequestList)
	require.Equal(t, genesisState.SignatureRequestCount, got.SignatureRequestCount)
	require.ElementsMatch(t, genesisState.SignatureShareList, got.SignatureShareList)
	require.Equal(t, genesisState.SignatureShareCount, got.SignatureShareCount)
	require.ElementsMatch(t, genesisState.SignedMessageList, got.SignedMessageList)
	require.Equal(t, genesisState.SignedMessageCount, got.SignedMessageCount)
	require.ElementsMatch(t, genesisState.MeditationSummoningList, got.MeditationSummoningList)
	require.Equal(t, genesisState.MeditationSummoningCount, got.MeditationSummoningCount)
	require.ElementsMatch(t, genesisState.MeditationList, got.MeditationList)
	require.Equal(t, genesisState.MeditationCount, got.MeditationCount)
	require.ElementsMatch(t, genesisState.ScriptureList, got.ScriptureList)
	require.ElementsMatch(t, genesisState.ScriptureSignatureRequestList, got.ScriptureSignatureRequestList)
	require.Equal(t, genesisState.ScriptureSignatureRequestCount, got.ScriptureSignatureRequestCount)
	require.ElementsMatch(t, genesisState.ScriptureSignatureShareList, got.ScriptureSignatureShareList)
	require.Equal(t, genesisState.ScriptureSignatureShareCount, got.ScriptureSignatureShareCount)
	require.ElementsMatch(t, genesisState.SignedScriptureList, got.SignedScriptureList)
	require.ElementsMatch(t, genesisState.SignedScriptureListList, got.SignedScriptureListList)
	require.Equal(t, genesisState.SignedScriptureListCount, got.SignedScriptureListCount)
	require.ElementsMatch(t, genesisState.BlessingList, got.BlessingList)
	require.Equal(t, genesisState.BlessingCount, got.BlessingCount)
	require.ElementsMatch(t, genesisState.BlessingReceiptList, got.BlessingReceiptList)
	require.ElementsMatch(t, genesisState.ImploringList, got.ImploringList)
	require.Equal(t, genesisState.ImploringCount, got.ImploringCount)
	require.ElementsMatch(t, genesisState.KillConjuringList, got.KillConjuringList)
	require.Equal(t, genesisState.KillConjuringCount, got.KillConjuringCount)
	require.ElementsMatch(t, genesisState.KillImploringList, got.KillImploringList)
	require.Equal(t, genesisState.KillImploringCount, got.KillImploringCount)
	require.ElementsMatch(t, genesisState.KillMeditationSummoningList, got.KillMeditationSummoningList)
	require.Equal(t, genesisState.KillMeditationSummoningCount, got.KillMeditationSummoningCount)
	require.ElementsMatch(t, genesisState.KillMagicKeySummoningList, got.KillMagicKeySummoningList)
	require.Equal(t, genesisState.KillMagicKeySummoningCount, got.KillMagicKeySummoningCount)
	require.ElementsMatch(t, genesisState.KillScriptureSignatureRequestList, got.KillScriptureSignatureRequestList)
	require.Equal(t, genesisState.KillScriptureSignatureRequestCount, got.KillScriptureSignatureRequestCount)
	require.ElementsMatch(t, genesisState.KillSignatureRequestList, got.KillSignatureRequestList)
	require.Equal(t, genesisState.KillSignatureRequestCount, got.KillSignatureRequestCount)
	require.ElementsMatch(t, genesisState.EncryptedMagicKeyShareList, got.EncryptedMagicKeyShareList)
	require.ElementsMatch(t, genesisState.EncryptedPreSignList, got.EncryptedPreSignList)
	// this line is used by starport scaffolding # genesis/test/assert
}
