package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMagicKey{}, "eightball/CreateMagicKey", nil)
	cdc.RegisterConcrete(&MsgCreateSpiritConjuringPoems{}, "eightball/CreateSpiritConjuringPoems", nil)
	cdc.RegisterConcrete(&MsgFinalizeConjuring{}, "eightball/FinalizeConjuring", nil)
	cdc.RegisterConcrete(&MsgCreateVerse{}, "eightball/CreateVerse", nil)
	cdc.RegisterConcrete(&MsgDeleteVerse{}, "eightball/DeleteVerse", nil)
	cdc.RegisterConcrete(&MsgCreateVision{}, "eightball/CreateVision", nil)
	cdc.RegisterConcrete(&MsgDeleteVision{}, "eightball/DeleteVision", nil)
	cdc.RegisterConcrete(&MsgCreateMessage{}, "eightball/CreateMessage", nil)
	cdc.RegisterConcrete(&MsgCreateSignatureRequest{}, "eightball/CreateSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureRequest{}, "eightball/UpdateSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteSignatureRequest{}, "eightball/DeleteSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgCreateSignatureShare{}, "eightball/CreateSignatureShare", nil)
	cdc.RegisterConcrete(&MsgCreateSignedMessage{}, "eightball/CreateSignedMessage", nil)
	cdc.RegisterConcrete(&MsgCreateMeditationSummoning{}, "eightball/CreateMeditationSummoning", nil)
	cdc.RegisterConcrete(&MsgDeleteMeditationSummoning{}, "eightball/DeleteMeditationSummoning", nil)
	cdc.RegisterConcrete(&MsgCreateMeditation{}, "eightball/CreateMeditation", nil)
	cdc.RegisterConcrete(&MsgDeleteMeditation{}, "eightball/DeleteMeditation", nil)
	cdc.RegisterConcrete(&MsgCreateScripture{}, "eightball/CreateScripture", nil)
	cdc.RegisterConcrete(&MsgCreateScriptureSignatureRequest{}, "eightball/CreateScriptureSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgCreateScriptureSignatureShare{}, "eightball/CreateScriptureSignatureShare", nil)
	cdc.RegisterConcrete(&MsgCreateSignedScripture{}, "eightball/CreateSignedScripture", nil)
	cdc.RegisterConcrete(&MsgCreateSignedScriptureList{}, "eightball/CreateSignedScriptureList", nil)
	cdc.RegisterConcrete(&MsgCreateBlessing{}, "eightball/CreateBlessing", nil)
	cdc.RegisterConcrete(&MsgCreateKillConjuring{}, "eightball/CreateKillConjuring", nil)
	cdc.RegisterConcrete(&MsgCreateKillImploring{}, "eightball/CreateKillImploring", nil)
	cdc.RegisterConcrete(&MsgCreateKillMeditationSummoning{}, "eightball/CreateKillMeditationSummoning", nil)
	cdc.RegisterConcrete(&MsgCreateKillMagicKeySummoning{}, "eightball/CreateKillMagicKeySummoning", nil)
	cdc.RegisterConcrete(&MsgCreateKillScriptureSignatureRequest{}, "eightball/CreateKillScriptureSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgCreateKillSignatureRequest{}, "eightball/CreateKillSignatureRequest", nil)
	cdc.RegisterConcrete(&MsgCreateEncryptedMagicKeyShare{}, "eightball/CreateEncryptedMagicKeyShare", nil)
	cdc.RegisterConcrete(&MsgDeleteEncryptedMagicKeyShare{}, "eightball/DeleteEncryptedMagicKeyShare", nil)
	cdc.RegisterConcrete(&MsgCreateEncryptedPreSign{}, "eightball/CreateEncryptedPreSign", nil)
	cdc.RegisterConcrete(&MsgDeleteEncryptedPreSign{}, "eightball/DeleteEncryptedPreSign", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMagicKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSpiritConjuringPoems{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalizeConjuring{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVerse{},
		&MsgDeleteVerse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVision{},
		&MsgDeleteVision{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMessage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignatureRequest{},
		&MsgUpdateSignatureRequest{},
		&MsgDeleteSignatureRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignatureShare{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignedMessage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMeditationSummoning{},
		&MsgDeleteMeditationSummoning{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMeditation{},
		&MsgDeleteMeditation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateScripture{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateScriptureSignatureRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateScriptureSignatureShare{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignedScripture{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignedScriptureList{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBlessing{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillConjuring{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillImploring{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillMeditationSummoning{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillMagicKeySummoning{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillScriptureSignatureRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKillSignatureRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEncryptedMagicKeyShare{},
		&MsgDeleteEncryptedMagicKeyShare{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEncryptedPreSign{},
		&MsgDeleteEncryptedPreSign{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
