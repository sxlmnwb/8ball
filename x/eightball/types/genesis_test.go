package types_test

import (
	"testing"

	"eightball/x/eightball/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				CurrentMagicKey: &types.CurrentMagicKey{
					CurrentMagicKey: 10,
					SinceBlock:      85,
				},
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated magicKey",
			genState: &types.GenesisState{
				MagicKeyList: []types.MagicKey{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid magicKey count",
			genState: &types.GenesisState{
				MagicKeyList: []types.MagicKey{
					{
						Id: 1,
					},
				},
				MagicKeyCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated magicKeySummoning",
			genState: &types.GenesisState{
				MagicKeySummoningList: []types.MagicKeySummoning{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid magicKeySummoning count",
			genState: &types.GenesisState{
				MagicKeySummoningList: []types.MagicKeySummoning{
					{
						Id: 1,
					},
				},
				MagicKeySummoningCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated highCouncil",
			genState: &types.GenesisState{
				HighCouncilList: []types.HighCouncil{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid highCouncil count",
			genState: &types.GenesisState{
				HighCouncilList: []types.HighCouncil{
					{
						Id: 1,
					},
				},
				HighCouncilCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated highCouncilConjurings",
			genState: &types.GenesisState{
				HighCouncilConjuringsList: []types.HighCouncilConjurings{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid highCouncilConjurings count",
			genState: &types.GenesisState{
				HighCouncilConjuringsList: []types.HighCouncilConjurings{
					{
						Id: 1,
					},
				},
				HighCouncilConjuringsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated spiritConjuringPoems",
			genState: &types.GenesisState{
				SpiritConjuringPoemsList: []types.SpiritConjuringPoems{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid spiritConjuringPoems count",
			genState: &types.GenesisState{
				SpiritConjuringPoemsList: []types.SpiritConjuringPoems{
					{
						Id: 1,
					},
				},
				SpiritConjuringPoemsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated verse",
			genState: &types.GenesisState{
				VerseList: []types.Verse{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid verse count",
			genState: &types.GenesisState{
				VerseList: []types.Verse{
					{
						Id: 1,
					},
				},
				VerseCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated vision",
			genState: &types.GenesisState{
				VisionList: []types.Vision{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid vision count",
			genState: &types.GenesisState{
				VisionList: []types.Vision{
					{
						Id: 1,
					},
				},
				VisionCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated message",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid message count",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 1,
					},
				},
				MessageCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated signatureRequest",
			genState: &types.GenesisState{
				SignatureRequestList: []types.SignatureRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid signatureRequest count",
			genState: &types.GenesisState{
				SignatureRequestList: []types.SignatureRequest{
					{
						Id: 1,
					},
				},
				SignatureRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated signatureShare",
			genState: &types.GenesisState{
				SignatureShareList: []types.SignatureShare{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid signatureShare count",
			genState: &types.GenesisState{
				SignatureShareList: []types.SignatureShare{
					{
						Id: 1,
					},
				},
				SignatureShareCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated signedMessage",
			genState: &types.GenesisState{
				SignedMessageList: []types.SignedMessage{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid signedMessage count",
			genState: &types.GenesisState{
				SignedMessageList: []types.SignedMessage{
					{
						Id: 1,
					},
				},
				SignedMessageCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated meditationSummoning",
			genState: &types.GenesisState{
				MeditationSummoningList: []types.MeditationSummoning{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid meditationSummoning count",
			genState: &types.GenesisState{
				MeditationSummoningList: []types.MeditationSummoning{
					{
						Id: 1,
					},
				},
				MeditationSummoningCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated meditation",
			genState: &types.GenesisState{
				MeditationList: []types.Meditation{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid meditation count",
			genState: &types.GenesisState{
				MeditationList: []types.Meditation{
					{
						Id: 1,
					},
				},
				MeditationCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated scripture",
			genState: &types.GenesisState{
				ScriptureList: []types.Scripture{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated scriptureSignatureRequest",
			genState: &types.GenesisState{
				ScriptureSignatureRequestList: []types.ScriptureSignatureRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid scriptureSignatureRequest count",
			genState: &types.GenesisState{
				ScriptureSignatureRequestList: []types.ScriptureSignatureRequest{
					{
						Id: 1,
					},
				},
				ScriptureSignatureRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated scriptureSignatureShare",
			genState: &types.GenesisState{
				ScriptureSignatureShareList: []types.ScriptureSignatureShare{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid scriptureSignatureShare count",
			genState: &types.GenesisState{
				ScriptureSignatureShareList: []types.ScriptureSignatureShare{
					{
						Id: 1,
					},
				},
				ScriptureSignatureShareCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated signedScripture",
			genState: &types.GenesisState{
				SignedScriptureList: []types.SignedScripture{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated signedScriptureList",
			genState: &types.GenesisState{
				SignedScriptureListList: []types.SignedScriptureList{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid signedScriptureList count",
			genState: &types.GenesisState{
				SignedScriptureListList: []types.SignedScriptureList{
					{
						Id: 1,
					},
				},
				SignedScriptureListCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated blessing",
			genState: &types.GenesisState{
				BlessingList: []types.Blessing{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid blessing count",
			genState: &types.GenesisState{
				BlessingList: []types.Blessing{
					{
						Id: 1,
					},
				},
				BlessingCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated blessingReceipt",
			genState: &types.GenesisState{
				BlessingReceiptList: []types.BlessingReceipt{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated imploring",
			genState: &types.GenesisState{
				ImploringList: []types.Imploring{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid imploring count",
			genState: &types.GenesisState{
				ImploringList: []types.Imploring{
					{
						Id: 1,
					},
				},
				ImploringCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killConjuring",
			genState: &types.GenesisState{
				KillConjuringList: []types.KillConjuring{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killConjuring count",
			genState: &types.GenesisState{
				KillConjuringList: []types.KillConjuring{
					{
						Id: 1,
					},
				},
				KillConjuringCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killImploring",
			genState: &types.GenesisState{
				KillImploringList: []types.KillImploring{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killImploring count",
			genState: &types.GenesisState{
				KillImploringList: []types.KillImploring{
					{
						Id: 1,
					},
				},
				KillImploringCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killMeditationSummoning",
			genState: &types.GenesisState{
				KillMeditationSummoningList: []types.KillMeditationSummoning{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killMeditationSummoning count",
			genState: &types.GenesisState{
				KillMeditationSummoningList: []types.KillMeditationSummoning{
					{
						Id: 1,
					},
				},
				KillMeditationSummoningCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killMagicKeySummoning",
			genState: &types.GenesisState{
				KillMagicKeySummoningList: []types.KillMagicKeySummoning{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killMagicKeySummoning count",
			genState: &types.GenesisState{
				KillMagicKeySummoningList: []types.KillMagicKeySummoning{
					{
						Id: 1,
					},
				},
				KillMagicKeySummoningCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killScriptureSignatureRequest",
			genState: &types.GenesisState{
				KillScriptureSignatureRequestList: []types.KillScriptureSignatureRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killScriptureSignatureRequest count",
			genState: &types.GenesisState{
				KillScriptureSignatureRequestList: []types.KillScriptureSignatureRequest{
					{
						Id: 1,
					},
				},
				KillScriptureSignatureRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated killSignatureRequest",
			genState: &types.GenesisState{
				KillSignatureRequestList: []types.KillSignatureRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid killSignatureRequest count",
			genState: &types.GenesisState{
				KillSignatureRequestList: []types.KillSignatureRequest{
					{
						Id: 1,
					},
				},
				KillSignatureRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated encryptedMagicKeyShare",
			genState: &types.GenesisState{
				EncryptedMagicKeyShareList: []types.EncryptedMagicKeyShare{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated encryptedPreSign",
			genState: &types.GenesisState{
				EncryptedPreSignList: []types.EncryptedPreSign{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
