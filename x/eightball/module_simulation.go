package eightball

import (
	"math/rand"

	"eightball/testutil/sample"
	eightballsimulation "eightball/x/eightball/simulation"
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = eightballsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateMagicKey = "op_weight_msg_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMagicKey int = 100

	opWeightMsgUpdateMagicKey = "op_weight_msg_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMagicKey int = 100

	opWeightMsgDeleteMagicKey = "op_weight_msg_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMagicKey int = 100

	opWeightMsgCreateCurrentMagicKey = "op_weight_msg_current_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCurrentMagicKey int = 100

	opWeightMsgUpdateCurrentMagicKey = "op_weight_msg_current_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCurrentMagicKey int = 100

	opWeightMsgDeleteCurrentMagicKey = "op_weight_msg_current_magic_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCurrentMagicKey int = 100

	opWeightMsgCreateSpiritConjuringPoems = "op_weight_msg_spirit_conjuring_poems"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSpiritConjuringPoems int = 100

	opWeightMsgUpdateSpiritConjuringPoems = "op_weight_msg_spirit_conjuring_poems"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSpiritConjuringPoems int = 100

	opWeightMsgDeleteSpiritConjuringPoems = "op_weight_msg_spirit_conjuring_poems"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSpiritConjuringPoems int = 100

	opWeightMsgFinalizeConjuring = "op_weight_msg_finalize_conjuring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinalizeConjuring int = 100

	opWeightMsgCreateVerse = "op_weight_msg_verse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVerse int = 100

	opWeightMsgUpdateVerse = "op_weight_msg_verse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVerse int = 100

	opWeightMsgDeleteVerse = "op_weight_msg_verse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVerse int = 100

	opWeightMsgCreateVision = "op_weight_msg_vision"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVision int = 100

	opWeightMsgUpdateVision = "op_weight_msg_vision"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVision int = 100

	opWeightMsgDeleteVision = "op_weight_msg_vision"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVision int = 100

	opWeightMsgCreateMessage = "op_weight_msg_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMessage int = 100

	opWeightMsgUpdateMessage = "op_weight_msg_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMessage int = 100

	opWeightMsgDeleteMessage = "op_weight_msg_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMessage int = 100

	opWeightMsgCreateSignatureRequest = "op_weight_msg_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignatureRequest int = 100

	opWeightMsgUpdateSignatureRequest = "op_weight_msg_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignatureRequest int = 100

	opWeightMsgDeleteSignatureRequest = "op_weight_msg_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignatureRequest int = 100

	opWeightMsgCreateSignatureShare = "op_weight_msg_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignatureShare int = 100

	opWeightMsgUpdateSignatureShare = "op_weight_msg_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignatureShare int = 100

	opWeightMsgDeleteSignatureShare = "op_weight_msg_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignatureShare int = 100

	opWeightMsgCreateSignedMessage = "op_weight_msg_signed_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignedMessage int = 100

	opWeightMsgUpdateSignedMessage = "op_weight_msg_signed_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignedMessage int = 100

	opWeightMsgDeleteSignedMessage = "op_weight_msg_signed_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignedMessage int = 100

	opWeightMsgCreateMeditationSummoning = "op_weight_msg_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMeditationSummoning int = 100

	opWeightMsgUpdateMeditationSummoning = "op_weight_msg_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMeditationSummoning int = 100

	opWeightMsgDeleteMeditationSummoning = "op_weight_msg_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMeditationSummoning int = 100

	opWeightMsgCreateMeditation = "op_weight_msg_meditation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMeditation int = 100

	opWeightMsgUpdateMeditation = "op_weight_msg_meditation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMeditation int = 100

	opWeightMsgDeleteMeditation = "op_weight_msg_meditation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMeditation int = 100

	opWeightMsgCreateScripture = "op_weight_msg_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateScripture int = 100

	opWeightMsgUpdateScripture = "op_weight_msg_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateScripture int = 100

	opWeightMsgDeleteScripture = "op_weight_msg_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteScripture int = 100

	opWeightMsgCreateScriptureSignatureRequest = "op_weight_msg_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateScriptureSignatureRequest int = 100

	opWeightMsgUpdateScriptureSignatureRequest = "op_weight_msg_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateScriptureSignatureRequest int = 100

	opWeightMsgDeleteScriptureSignatureRequest = "op_weight_msg_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteScriptureSignatureRequest int = 100

	opWeightMsgCreateScriptureSignatureShare = "op_weight_msg_scripture_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateScriptureSignatureShare int = 100

	opWeightMsgUpdateScriptureSignatureShare = "op_weight_msg_scripture_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateScriptureSignatureShare int = 100

	opWeightMsgDeleteScriptureSignatureShare = "op_weight_msg_scripture_signature_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteScriptureSignatureShare int = 100

	opWeightMsgCreateSignedScripture = "op_weight_msg_signed_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignedScripture int = 100

	opWeightMsgUpdateSignedScripture = "op_weight_msg_signed_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignedScripture int = 100

	opWeightMsgDeleteSignedScripture = "op_weight_msg_signed_scripture"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignedScripture int = 100

	opWeightMsgCreateSignedScriptureList = "op_weight_msg_signed_scripture_list"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSignedScriptureList int = 100

	opWeightMsgUpdateSignedScriptureList = "op_weight_msg_signed_scripture_list"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSignedScriptureList int = 100

	opWeightMsgDeleteSignedScriptureList = "op_weight_msg_signed_scripture_list"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSignedScriptureList int = 100

	opWeightMsgCreateBlessing = "op_weight_msg_blessing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBlessing int = 100

	opWeightMsgUpdateBlessing = "op_weight_msg_blessing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBlessing int = 100

	opWeightMsgDeleteBlessing = "op_weight_msg_blessing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBlessing int = 100

	opWeightMsgCreateKillConjuring = "op_weight_msg_kill_conjuring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillConjuring int = 100

	opWeightMsgUpdateKillConjuring = "op_weight_msg_kill_conjuring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillConjuring int = 100

	opWeightMsgDeleteKillConjuring = "op_weight_msg_kill_conjuring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillConjuring int = 100

	opWeightMsgCreateKillImploring = "op_weight_msg_kill_imploring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillImploring int = 100

	opWeightMsgUpdateKillImploring = "op_weight_msg_kill_imploring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillImploring int = 100

	opWeightMsgDeleteKillImploring = "op_weight_msg_kill_imploring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillImploring int = 100

	opWeightMsgCreateKillMeditationSummoning = "op_weight_msg_kill_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillMeditationSummoning int = 100

	opWeightMsgUpdateKillMeditationSummoning = "op_weight_msg_kill_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillMeditationSummoning int = 100

	opWeightMsgDeleteKillMeditationSummoning = "op_weight_msg_kill_meditation_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillMeditationSummoning int = 100

	opWeightMsgCreateKillMagicKeySummoning = "op_weight_msg_kill_magic_key_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillMagicKeySummoning int = 100

	opWeightMsgUpdateKillMagicKeySummoning = "op_weight_msg_kill_magic_key_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillMagicKeySummoning int = 100

	opWeightMsgDeleteKillMagicKeySummoning = "op_weight_msg_kill_magic_key_summoning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillMagicKeySummoning int = 100

	opWeightMsgCreateKillScriptureSignatureRequest = "op_weight_msg_kill_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillScriptureSignatureRequest int = 100

	opWeightMsgUpdateKillScriptureSignatureRequest = "op_weight_msg_kill_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillScriptureSignatureRequest int = 100

	opWeightMsgDeleteKillScriptureSignatureRequest = "op_weight_msg_kill_scripture_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillScriptureSignatureRequest int = 100

	opWeightMsgCreateKillSignatureRequest = "op_weight_msg_kill_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKillSignatureRequest int = 100

	opWeightMsgUpdateKillSignatureRequest = "op_weight_msg_kill_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKillSignatureRequest int = 100

	opWeightMsgDeleteKillSignatureRequest = "op_weight_msg_kill_signature_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKillSignatureRequest int = 100

	opWeightMsgCreateEncryptedMagicKeyShare = "op_weight_msg_encrypted_magic_key_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEncryptedMagicKeyShare int = 100

	opWeightMsgUpdateEncryptedMagicKeyShare = "op_weight_msg_encrypted_magic_key_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEncryptedMagicKeyShare int = 100

	opWeightMsgDeleteEncryptedMagicKeyShare = "op_weight_msg_encrypted_magic_key_share"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEncryptedMagicKeyShare int = 100

	opWeightMsgCreateEncryptedPreSign = "op_weight_msg_encrypted_pre_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEncryptedPreSign int = 100

	opWeightMsgUpdateEncryptedPreSign = "op_weight_msg_encrypted_pre_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEncryptedPreSign int = 100

	opWeightMsgDeleteEncryptedPreSign = "op_weight_msg_encrypted_pre_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEncryptedPreSign int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	eightballGenesis := types.GenesisState{
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
		SpiritConjuringPoemsList: []types.SpiritConjuringPoems{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SpiritConjuringPoemsCount: 2,
		VerseList: []types.Verse{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		VerseCount: 2,
		VisionList: []types.Vision{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		VisionCount: 2,
		MessageList: []types.Message{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MessageCount: 2,
		SignatureRequestList: []types.SignatureRequest{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SignatureRequestCount: 2,
		SignatureShareList: []types.SignatureShare{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SignatureShareCount: 2,
		SignedMessageList: []types.SignedMessage{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SignedMessageCount: 2,
		MeditationSummoningList: []types.MeditationSummoning{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MeditationSummoningCount: 2,
		MeditationList: []types.Meditation{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
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
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
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
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SignedScriptureListCount: 2,
		BlessingList: []types.Blessing{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		BlessingCount: 2,
		KillConjuringList: []types.KillConjuring{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillConjuringCount: 2,
		KillImploringList: []types.KillImploring{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillImploringCount: 2,
		KillMeditationSummoningList: []types.KillMeditationSummoning{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillMeditationSummoningCount: 2,
		KillMagicKeySummoningList: []types.KillMagicKeySummoning{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillMagicKeySummoningCount: 2,
		KillScriptureSignatureRequestList: []types.KillScriptureSignatureRequest{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillScriptureSignatureRequestCount: 2,
		KillSignatureRequestList: []types.KillSignatureRequest{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KillSignatureRequestCount: 2,
		EncryptedMagicKeyShareList: []types.EncryptedMagicKeyShare{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		EncryptedPreSignList: []types.EncryptedPreSign{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&eightballGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMagicKey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMagicKey, &weightMsgCreateMagicKey, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMagicKey = defaultWeightMsgCreateMagicKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMagicKey,
		eightballsimulation.SimulateMsgCreateMagicKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMagicKey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMagicKey, &weightMsgUpdateMagicKey, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMagicKey = defaultWeightMsgUpdateMagicKey
		},
	)

	var weightMsgDeleteMagicKey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMagicKey, &weightMsgDeleteMagicKey, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMagicKey = defaultWeightMsgDeleteMagicKey
		},
	)

	var weightMsgCreateSpiritConjuringPoems int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSpiritConjuringPoems, &weightMsgCreateSpiritConjuringPoems, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSpiritConjuringPoems = defaultWeightMsgCreateSpiritConjuringPoems
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSpiritConjuringPoems,
		eightballsimulation.SimulateMsgCreateSpiritConjuringPoems(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSpiritConjuringPoems int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSpiritConjuringPoems, &weightMsgUpdateSpiritConjuringPoems, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSpiritConjuringPoems = defaultWeightMsgUpdateSpiritConjuringPoems
		},
	)

	var weightMsgDeleteSpiritConjuringPoems int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSpiritConjuringPoems, &weightMsgDeleteSpiritConjuringPoems, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSpiritConjuringPoems = defaultWeightMsgDeleteSpiritConjuringPoems
		},
	)

	var weightMsgFinalizeConjuring int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinalizeConjuring, &weightMsgFinalizeConjuring, nil,
		func(_ *rand.Rand) {
			weightMsgFinalizeConjuring = defaultWeightMsgFinalizeConjuring
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinalizeConjuring,
		eightballsimulation.SimulateMsgFinalizeConjuring(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVerse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVerse, &weightMsgCreateVerse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVerse = defaultWeightMsgCreateVerse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVerse,
		eightballsimulation.SimulateMsgCreateVerse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVerse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVerse, &weightMsgUpdateVerse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVerse = defaultWeightMsgUpdateVerse
		},
	)

	var weightMsgDeleteVerse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVerse, &weightMsgDeleteVerse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVerse = defaultWeightMsgDeleteVerse
		},
	)

	var weightMsgCreateVision int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVision, &weightMsgCreateVision, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVision = defaultWeightMsgCreateVision
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVision,
		eightballsimulation.SimulateMsgCreateVision(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVision int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVision, &weightMsgUpdateVision, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVision = defaultWeightMsgUpdateVision
		},
	)

	var weightMsgDeleteVision int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVision, &weightMsgDeleteVision, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVision = defaultWeightMsgDeleteVision
		},
	)
	var weightMsgCreateMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMessage, &weightMsgCreateMessage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMessage = defaultWeightMsgCreateMessage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMessage,
		eightballsimulation.SimulateMsgCreateMessage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMessage, &weightMsgUpdateMessage, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMessage = defaultWeightMsgUpdateMessage
		},
	)

	var weightMsgDeleteMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMessage, &weightMsgDeleteMessage, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMessage = defaultWeightMsgDeleteMessage
		},
	)

	var weightMsgCreateSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignatureRequest, &weightMsgCreateSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignatureRequest = defaultWeightMsgCreateSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignatureRequest,
		eightballsimulation.SimulateMsgCreateSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSignatureRequest, &weightMsgUpdateSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSignatureRequest = defaultWeightMsgUpdateSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSignatureRequest,
		eightballsimulation.SimulateMsgUpdateSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSignatureRequest, &weightMsgDeleteSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSignatureRequest = defaultWeightMsgDeleteSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSignatureRequest,
		eightballsimulation.SimulateMsgDeleteSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSignatureShare int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignatureShare, &weightMsgCreateSignatureShare, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignatureShare = defaultWeightMsgCreateSignatureShare
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignatureShare,
		eightballsimulation.SimulateMsgCreateSignatureShare(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSignatureShare int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSignatureShare, &weightMsgUpdateSignatureShare, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSignatureShare = defaultWeightMsgUpdateSignatureShare
		},
	)

	var weightMsgDeleteSignatureShare int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSignatureShare, &weightMsgDeleteSignatureShare, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSignatureShare = defaultWeightMsgDeleteSignatureShare
		},
	)

	var weightMsgCreateSignedMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignedMessage, &weightMsgCreateSignedMessage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignedMessage = defaultWeightMsgCreateSignedMessage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignedMessage,
		eightballsimulation.SimulateMsgCreateSignedMessage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSignedMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSignedMessage, &weightMsgUpdateSignedMessage, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSignedMessage = defaultWeightMsgUpdateSignedMessage
		},
	)

	var weightMsgDeleteSignedMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSignedMessage, &weightMsgDeleteSignedMessage, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSignedMessage = defaultWeightMsgDeleteSignedMessage
		},
	)

	var weightMsgCreateMeditationSummoning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMeditationSummoning, &weightMsgCreateMeditationSummoning, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMeditationSummoning = defaultWeightMsgCreateMeditationSummoning
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMeditationSummoning,
		eightballsimulation.SimulateMsgCreateMeditationSummoning(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMeditationSummoning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMeditationSummoning, &weightMsgUpdateMeditationSummoning, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMeditationSummoning = defaultWeightMsgUpdateMeditationSummoning
		},
	)

	var weightMsgDeleteMeditationSummoning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMeditationSummoning, &weightMsgDeleteMeditationSummoning, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMeditationSummoning = defaultWeightMsgDeleteMeditationSummoning
		},
	)

	var weightMsgCreateMeditation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMeditation, &weightMsgCreateMeditation, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMeditation = defaultWeightMsgCreateMeditation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMeditation,
		eightballsimulation.SimulateMsgCreateMeditation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMeditation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMeditation, &weightMsgUpdateMeditation, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMeditation = defaultWeightMsgUpdateMeditation
		},
	)

	var weightMsgDeleteMeditation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMeditation, &weightMsgDeleteMeditation, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMeditation = defaultWeightMsgDeleteMeditation
		},
	)
	var weightMsgCreateScripture int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateScripture, &weightMsgCreateScripture, nil,
		func(_ *rand.Rand) {
			weightMsgCreateScripture = defaultWeightMsgCreateScripture
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateScripture,
		eightballsimulation.SimulateMsgCreateScripture(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateScriptureSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateScriptureSignatureRequest, &weightMsgCreateScriptureSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateScriptureSignatureRequest = defaultWeightMsgCreateScriptureSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateScriptureSignatureRequest,
		eightballsimulation.SimulateMsgCreateScriptureSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateScriptureSignatureShare int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateScriptureSignatureShare, &weightMsgCreateScriptureSignatureShare, nil,
		func(_ *rand.Rand) {
			weightMsgCreateScriptureSignatureShare = defaultWeightMsgCreateScriptureSignatureShare
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateScriptureSignatureShare,
		eightballsimulation.SimulateMsgCreateScriptureSignatureShare(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSignedScripture int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignedScripture, &weightMsgCreateSignedScripture, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignedScripture = defaultWeightMsgCreateSignedScripture
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignedScripture,
		eightballsimulation.SimulateMsgCreateSignedScripture(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSignedScriptureList int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSignedScriptureList, &weightMsgCreateSignedScriptureList, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSignedScriptureList = defaultWeightMsgCreateSignedScriptureList
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSignedScriptureList,
		eightballsimulation.SimulateMsgCreateSignedScriptureList(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBlessing int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBlessing, &weightMsgCreateBlessing, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBlessing = defaultWeightMsgCreateBlessing
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBlessing,
		eightballsimulation.SimulateMsgCreateBlessing(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillConjuring int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillConjuring, &weightMsgCreateKillConjuring, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillConjuring = defaultWeightMsgCreateKillConjuring
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillConjuring,
		eightballsimulation.SimulateMsgCreateKillConjuring(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillImploring int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillImploring, &weightMsgCreateKillImploring, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillImploring = defaultWeightMsgCreateKillImploring
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillImploring,
		eightballsimulation.SimulateMsgCreateKillImploring(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillMeditationSummoning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillMeditationSummoning, &weightMsgCreateKillMeditationSummoning, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillMeditationSummoning = defaultWeightMsgCreateKillMeditationSummoning
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillMeditationSummoning,
		eightballsimulation.SimulateMsgCreateKillMeditationSummoning(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillMagicKeySummoning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillMagicKeySummoning, &weightMsgCreateKillMagicKeySummoning, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillMagicKeySummoning = defaultWeightMsgCreateKillMagicKeySummoning
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillMagicKeySummoning,
		eightballsimulation.SimulateMsgCreateKillMagicKeySummoning(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillScriptureSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillScriptureSignatureRequest, &weightMsgCreateKillScriptureSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillScriptureSignatureRequest = defaultWeightMsgCreateKillScriptureSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillScriptureSignatureRequest,
		eightballsimulation.SimulateMsgCreateKillScriptureSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKillSignatureRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKillSignatureRequest, &weightMsgCreateKillSignatureRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKillSignatureRequest = defaultWeightMsgCreateKillSignatureRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKillSignatureRequest,
		eightballsimulation.SimulateMsgCreateKillSignatureRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
