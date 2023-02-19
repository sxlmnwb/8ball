package types

const (
	// ModuleName defines the module name
	ModuleName = "eightball"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_eightball"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MagicKeyKey      = "MagicKey/value/"
	MagicKeyCountKey = "MagicKey/count/"
)

const (
	MagicKeySummoningKey      = "MagicKeySummoning/value/"
	MagicKeySummoningCountKey = "MagicKeySummoning/count/"
)

const (
	HighCouncilKey      = "HighCouncil/value/"
	HighCouncilCountKey = "HighCouncil/count/"
)

const (
	CurrentMagicKeyKey = "CurrentMagicKey/value/"
)

const (
	HighCouncilConjuringsKey      = "HighCouncilConjurings/value/"
	HighCouncilConjuringsCountKey = "HighCouncilConjurings/count/"
)

const (
	SpiritConjuringPoemsKey      = "SpiritConjuringPoems/value/"
	SpiritConjuringPoemsCountKey = "SpiritConjuringPoems/count/"
)

const (
	VerseKey      = "Verse/value/"
	VerseCountKey = "Verse/count/"
)

const (
	VisionKey      = "Vision/value/"
	VisionCountKey = "Vision/count/"
)

const (
	MessageKey      = "Message/value/"
	MessageCountKey = "Message/count/"
)

const (
	SignatureRequestKey      = "SignatureRequest/value/"
	SignatureRequestCountKey = "SignatureRequest/count/"
)

const (
	SignatureShareKey      = "SignatureShare/value/"
	SignatureShareCountKey = "SignatureShare/count/"
)

const (
	SignedMessageKey      = "SignedMessage/value/"
	SignedMessageCountKey = "SignedMessage/count/"
)

const (
	MeditationSummoningKey      = "MeditationSummoning/value/"
	MeditationSummoningCountKey = "MeditationSummoning/count/"
)

const (
	MeditationKey      = "Meditation/value/"
	MeditationCountKey = "Meditation/count/"
)

const (
	ScriptureSignatureRequestKey      = "ScriptureSignatureRequest/value/"
	ScriptureSignatureRequestCountKey = "ScriptureSignatureRequest/count/"
)

const (
	ScriptureSignatureShareKey      = "ScriptureSignatureShare/value/"
	ScriptureSignatureShareCountKey = "ScriptureSignatureShare/count/"
)

const (
	SignedScriptureListKey      = "SignedScriptureList/value/"
	SignedScriptureListCountKey = "SignedScriptureList/count/"
)

const (
	BlessingKey      = "Blessing/value/"
	BlessingCountKey = "Blessing/count/"
)

const (
	ImploringKey      = "Imploring/value/"
	ImploringCountKey = "Imploring/count/"
)

const (
	KillConjuringKey      = "KillConjuring/value/"
	KillConjuringCountKey = "KillConjuring/count/"
)

const (
	KillImploringKey      = "KillImploring/value/"
	KillImploringCountKey = "KillImploring/count/"
)

const (
	KillMeditationSummoningKey      = "KillMeditationSummoning/value/"
	KillMeditationSummoningCountKey = "KillMeditationSummoning/count/"
)

const (
	KillMagicKeySummoningKey      = "KillMagicKeySummoning/value/"
	KillMagicKeySummoningCountKey = "KillMagicKeySummoning/count/"
)

const (
	KillScriptureSignatureRequestKey      = "KillScriptureSignatureRequest/value/"
	KillScriptureSignatureRequestCountKey = "KillScriptureSignatureRequest/count/"
)

const (
	KillSignatureRequestKey      = "KillSignatureRequest/value/"
	KillSignatureRequestCountKey = "KillSignatureRequest/count/"
)
