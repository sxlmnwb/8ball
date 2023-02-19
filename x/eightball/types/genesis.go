package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MagicKeyList:                      []MagicKey{},
		MagicKeySummoningList:             []MagicKeySummoning{},
		HighCouncilList:                   []HighCouncil{},
		CurrentMagicKey:                   nil,
		HighCouncilConjuringsList:         []HighCouncilConjurings{},
		SpiritConjuringPoemsList:          []SpiritConjuringPoems{},
		VerseList:                         []Verse{},
		VisionList:                        []Vision{},
		MessageList:                       []Message{},
		SignatureRequestList:              []SignatureRequest{},
		SignatureShareList:                []SignatureShare{},
		SignedMessageList:                 []SignedMessage{},
		MeditationSummoningList:           []MeditationSummoning{},
		MeditationList:                    []Meditation{},
		ScriptureList:                     []Scripture{},
		ScriptureSignatureRequestList:     []ScriptureSignatureRequest{},
		ScriptureSignatureShareList:       []ScriptureSignatureShare{},
		SignedScriptureList:               []SignedScripture{},
		SignedScriptureListList:           []SignedScriptureList{},
		BlessingList:                      []Blessing{},
		BlessingReceiptList:               []BlessingReceipt{},
		ImploringList:                     []Imploring{},
		KillConjuringList:                 []KillConjuring{},
		KillImploringList:                 []KillImploring{},
		KillMeditationSummoningList:       []KillMeditationSummoning{},
		KillMagicKeySummoningList:         []KillMagicKeySummoning{},
		KillScriptureSignatureRequestList: []KillScriptureSignatureRequest{},
		KillSignatureRequestList:          []KillSignatureRequest{},
		EncryptedMagicKeyShareList:        []EncryptedMagicKeyShare{},
		EncryptedPreSignList:              []EncryptedPreSign{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in magicKey
	magicKeyIdMap := make(map[uint64]bool)
	magicKeyCount := gs.GetMagicKeyCount()
	for _, elem := range gs.MagicKeyList {
		if _, ok := magicKeyIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for magicKey")
		}
		if elem.Id >= magicKeyCount {
			return fmt.Errorf("magicKey id should be lower or equal than the last id")
		}
		magicKeyIdMap[elem.Id] = true
	}
	// Check for duplicated ID in magicKeySummoning
	magicKeySummoningIdMap := make(map[uint64]bool)
	magicKeySummoningCount := gs.GetMagicKeySummoningCount()
	for _, elem := range gs.MagicKeySummoningList {
		if _, ok := magicKeySummoningIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for magicKeySummoning")
		}
		if elem.Id >= magicKeySummoningCount {
			return fmt.Errorf("magicKeySummoning id should be lower or equal than the last id")
		}
		magicKeySummoningIdMap[elem.Id] = true
	}
	// Check for duplicated ID in highCouncil
	highCouncilIdMap := make(map[uint64]bool)
	highCouncilCount := gs.GetHighCouncilCount()
	for _, elem := range gs.HighCouncilList {
		if _, ok := highCouncilIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for highCouncil")
		}
		if elem.Id >= highCouncilCount {
			return fmt.Errorf("highCouncil id should be lower or equal than the last id")
		}
		highCouncilIdMap[elem.Id] = true
	}
	// Check for duplicated ID in highCouncilConjurings
	highCouncilConjuringsIdMap := make(map[uint64]bool)
	highCouncilConjuringsCount := gs.GetHighCouncilConjuringsCount()
	for _, elem := range gs.HighCouncilConjuringsList {
		if _, ok := highCouncilConjuringsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for highCouncilConjurings")
		}
		if elem.Id >= highCouncilConjuringsCount {
			return fmt.Errorf("highCouncilConjurings id should be lower or equal than the last id")
		}
		highCouncilConjuringsIdMap[elem.Id] = true
	}
	// Check for duplicated ID in spiritConjuringPoems
	spiritConjuringPoemsIdMap := make(map[uint64]bool)
	spiritConjuringPoemsCount := gs.GetSpiritConjuringPoemsCount()
	for _, elem := range gs.SpiritConjuringPoemsList {
		if _, ok := spiritConjuringPoemsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for spiritConjuringPoems")
		}
		if elem.Id >= spiritConjuringPoemsCount {
			return fmt.Errorf("spiritConjuringPoems id should be lower or equal than the last id")
		}
		spiritConjuringPoemsIdMap[elem.Id] = true
	}
	// Check for duplicated ID in verse
	verseIdMap := make(map[uint64]bool)
	verseCount := gs.GetVerseCount()
	for _, elem := range gs.VerseList {
		if _, ok := verseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for verse")
		}
		if elem.Id >= verseCount {
			return fmt.Errorf("verse id should be lower or equal than the last id")
		}
		verseIdMap[elem.Id] = true
	}
	// Check for duplicated ID in vision
	visionIdMap := make(map[uint64]bool)
	visionCount := gs.GetVisionCount()
	for _, elem := range gs.VisionList {
		if _, ok := visionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vision")
		}
		if elem.Id >= visionCount {
			return fmt.Errorf("vision id should be lower or equal than the last id")
		}
		visionIdMap[elem.Id] = true
	}
	// Check for duplicated ID in message
	messageIdMap := make(map[uint64]bool)
	messageCount := gs.GetMessageCount()
	for _, elem := range gs.MessageList {
		if _, ok := messageIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for message")
		}
		if elem.Id >= messageCount {
			return fmt.Errorf("message id should be lower or equal than the last id")
		}
		messageIdMap[elem.Id] = true
	}
	// Check for duplicated ID in signatureRequest
	signatureRequestIdMap := make(map[uint64]bool)
	signatureRequestCount := gs.GetSignatureRequestCount()
	for _, elem := range gs.SignatureRequestList {
		if _, ok := signatureRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for signatureRequest")
		}
		if elem.Id >= signatureRequestCount {
			return fmt.Errorf("signatureRequest id should be lower or equal than the last id")
		}
		signatureRequestIdMap[elem.Id] = true
	}
	// Check for duplicated ID in signatureShare
	signatureShareIdMap := make(map[uint64]bool)
	signatureShareCount := gs.GetSignatureShareCount()
	for _, elem := range gs.SignatureShareList {
		if _, ok := signatureShareIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for signatureShare")
		}
		if elem.Id >= signatureShareCount {
			return fmt.Errorf("signatureShare id should be lower or equal than the last id")
		}
		signatureShareIdMap[elem.Id] = true
	}
	// Check for duplicated ID in signedMessage
	signedMessageIdMap := make(map[uint64]bool)
	signedMessageCount := gs.GetSignedMessageCount()
	for _, elem := range gs.SignedMessageList {
		if _, ok := signedMessageIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for signedMessage")
		}
		if elem.Id >= signedMessageCount {
			return fmt.Errorf("signedMessage id should be lower or equal than the last id")
		}
		signedMessageIdMap[elem.Id] = true
	}
	// Check for duplicated ID in meditationSummoning
	meditationSummoningIdMap := make(map[uint64]bool)
	meditationSummoningCount := gs.GetMeditationSummoningCount()
	for _, elem := range gs.MeditationSummoningList {
		if _, ok := meditationSummoningIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for meditationSummoning")
		}
		if elem.Id >= meditationSummoningCount {
			return fmt.Errorf("meditationSummoning id should be lower or equal than the last id")
		}
		meditationSummoningIdMap[elem.Id] = true
	}
	// Check for duplicated ID in meditation
	meditationIdMap := make(map[uint64]bool)
	meditationCount := gs.GetMeditationCount()
	for _, elem := range gs.MeditationList {
		if _, ok := meditationIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for meditation")
		}
		if elem.Id >= meditationCount {
			return fmt.Errorf("meditation id should be lower or equal than the last id")
		}
		meditationIdMap[elem.Id] = true
	}
	// Check for duplicated index in scripture
	scriptureIndexMap := make(map[string]struct{})

	for _, elem := range gs.ScriptureList {
		index := string(ScriptureKey(elem.Index))
		if _, ok := scriptureIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for scripture")
		}
		scriptureIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in scriptureSignatureRequest
	scriptureSignatureRequestIdMap := make(map[uint64]bool)
	scriptureSignatureRequestCount := gs.GetScriptureSignatureRequestCount()
	for _, elem := range gs.ScriptureSignatureRequestList {
		if _, ok := scriptureSignatureRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for scriptureSignatureRequest")
		}
		if elem.Id >= scriptureSignatureRequestCount {
			return fmt.Errorf("scriptureSignatureRequest id should be lower or equal than the last id")
		}
		scriptureSignatureRequestIdMap[elem.Id] = true
	}
	// Check for duplicated ID in scriptureSignatureShare
	scriptureSignatureShareIdMap := make(map[uint64]bool)
	scriptureSignatureShareCount := gs.GetScriptureSignatureShareCount()
	for _, elem := range gs.ScriptureSignatureShareList {
		if _, ok := scriptureSignatureShareIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for scriptureSignatureShare")
		}
		if elem.Id >= scriptureSignatureShareCount {
			return fmt.Errorf("scriptureSignatureShare id should be lower or equal than the last id")
		}
		scriptureSignatureShareIdMap[elem.Id] = true
	}
	// Check for duplicated index in signedScripture
	signedScriptureIndexMap := make(map[string]struct{})

	for _, elem := range gs.SignedScriptureList {
		index := string(SignedScriptureKey(elem.Index))
		if _, ok := signedScriptureIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for signedScripture")
		}
		signedScriptureIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in signedScriptureList
	signedScriptureListIdMap := make(map[uint64]bool)
	signedScriptureListCount := gs.GetSignedScriptureListCount()
	for _, elem := range gs.SignedScriptureListList {
		if _, ok := signedScriptureListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for signedScriptureList")
		}
		if elem.Id >= signedScriptureListCount {
			return fmt.Errorf("signedScriptureList id should be lower or equal than the last id")
		}
		signedScriptureListIdMap[elem.Id] = true
	}
	// Check for duplicated ID in blessing
	blessingIdMap := make(map[uint64]bool)
	blessingCount := gs.GetBlessingCount()
	for _, elem := range gs.BlessingList {
		if _, ok := blessingIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for blessing")
		}
		if elem.Id >= blessingCount {
			return fmt.Errorf("blessing id should be lower or equal than the last id")
		}
		blessingIdMap[elem.Id] = true
	}
	// Check for duplicated index in blessingReceipt
	blessingReceiptIndexMap := make(map[string]struct{})

	for _, elem := range gs.BlessingReceiptList {
		index := string(BlessingReceiptKey(elem.Index))
		if _, ok := blessingReceiptIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for blessingReceipt")
		}
		blessingReceiptIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in imploring
	imploringIdMap := make(map[uint64]bool)
	imploringCount := gs.GetImploringCount()
	for _, elem := range gs.ImploringList {
		if _, ok := imploringIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for imploring")
		}
		if elem.Id >= imploringCount {
			return fmt.Errorf("imploring id should be lower or equal than the last id")
		}
		imploringIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killConjuring
	killConjuringIdMap := make(map[uint64]bool)
	killConjuringCount := gs.GetKillConjuringCount()
	for _, elem := range gs.KillConjuringList {
		if _, ok := killConjuringIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killConjuring")
		}
		if elem.Id >= killConjuringCount {
			return fmt.Errorf("killConjuring id should be lower or equal than the last id")
		}
		killConjuringIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killImploring
	killImploringIdMap := make(map[uint64]bool)
	killImploringCount := gs.GetKillImploringCount()
	for _, elem := range gs.KillImploringList {
		if _, ok := killImploringIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killImploring")
		}
		if elem.Id >= killImploringCount {
			return fmt.Errorf("killImploring id should be lower or equal than the last id")
		}
		killImploringIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killMeditationSummoning
	killMeditationSummoningIdMap := make(map[uint64]bool)
	killMeditationSummoningCount := gs.GetKillMeditationSummoningCount()
	for _, elem := range gs.KillMeditationSummoningList {
		if _, ok := killMeditationSummoningIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killMeditationSummoning")
		}
		if elem.Id >= killMeditationSummoningCount {
			return fmt.Errorf("killMeditationSummoning id should be lower or equal than the last id")
		}
		killMeditationSummoningIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killMagicKeySummoning
	killMagicKeySummoningIdMap := make(map[uint64]bool)
	killMagicKeySummoningCount := gs.GetKillMagicKeySummoningCount()
	for _, elem := range gs.KillMagicKeySummoningList {
		if _, ok := killMagicKeySummoningIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killMagicKeySummoning")
		}
		if elem.Id >= killMagicKeySummoningCount {
			return fmt.Errorf("killMagicKeySummoning id should be lower or equal than the last id")
		}
		killMagicKeySummoningIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killScriptureSignatureRequest
	killScriptureSignatureRequestIdMap := make(map[uint64]bool)
	killScriptureSignatureRequestCount := gs.GetKillScriptureSignatureRequestCount()
	for _, elem := range gs.KillScriptureSignatureRequestList {
		if _, ok := killScriptureSignatureRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killScriptureSignatureRequest")
		}
		if elem.Id >= killScriptureSignatureRequestCount {
			return fmt.Errorf("killScriptureSignatureRequest id should be lower or equal than the last id")
		}
		killScriptureSignatureRequestIdMap[elem.Id] = true
	}
	// Check for duplicated ID in killSignatureRequest
	killSignatureRequestIdMap := make(map[uint64]bool)
	killSignatureRequestCount := gs.GetKillSignatureRequestCount()
	for _, elem := range gs.KillSignatureRequestList {
		if _, ok := killSignatureRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for killSignatureRequest")
		}
		if elem.Id >= killSignatureRequestCount {
			return fmt.Errorf("killSignatureRequest id should be lower or equal than the last id")
		}
		killSignatureRequestIdMap[elem.Id] = true
	}
	// Check for duplicated index in encryptedMagicKeyShare
	encryptedMagicKeyShareIndexMap := make(map[string]struct{})

	for _, elem := range gs.EncryptedMagicKeyShareList {
		index := string(EncryptedMagicKeyShareKey(elem.Index))
		if _, ok := encryptedMagicKeyShareIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for encryptedMagicKeyShare")
		}
		encryptedMagicKeyShareIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in encryptedPreSign
	encryptedPreSignIndexMap := make(map[string]struct{})

	for _, elem := range gs.EncryptedPreSignList {
		index := string(EncryptedPreSignKey(elem.Index))
		if _, ok := encryptedPreSignIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for encryptedPreSign")
		}
		encryptedPreSignIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
