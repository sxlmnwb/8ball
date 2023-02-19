package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group eightball queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "8ball",
		Short:                      fmt.Sprintf("Querying commands for the %s module", "8ball"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
		Aliases:                    []string{"8ball"},
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListMagicKey())
	cmd.AddCommand(CmdShowMagicKey())
	cmd.AddCommand(CmdListMagicKeySummoning())
	cmd.AddCommand(CmdShowMagicKeySummoning())
	cmd.AddCommand(CmdListHighCouncil())
	cmd.AddCommand(CmdShowHighCouncil())
	cmd.AddCommand(CmdShowCurrentMagicKey())
	cmd.AddCommand(CmdListHighCouncilConjurings())
	cmd.AddCommand(CmdShowHighCouncilConjurings())
	cmd.AddCommand(CmdListSpiritConjuringPoems())
	cmd.AddCommand(CmdShowSpiritConjuringPoems())
	cmd.AddCommand(CmdListVerse())
	cmd.AddCommand(CmdShowVerse())
	cmd.AddCommand(CmdVerseCount())

	cmd.AddCommand(CmdMagicKeyCount())

	cmd.AddCommand(CmdListVision())
	cmd.AddCommand(CmdShowVision())
	cmd.AddCommand(CmdListMessage())
	cmd.AddCommand(CmdShowMessage())
	cmd.AddCommand(CmdListSignatureRequest())
	cmd.AddCommand(CmdShowSignatureRequest())
	cmd.AddCommand(CmdListSignatureShare())
	cmd.AddCommand(CmdShowSignatureShare())
	cmd.AddCommand(CmdListSignedMessage())
	cmd.AddCommand(CmdShowSignedMessage())
	cmd.AddCommand(CmdListMeditationSummoning())
	cmd.AddCommand(CmdShowMeditationSummoning())
	cmd.AddCommand(CmdListMeditation())
	cmd.AddCommand(CmdShowMeditation())
	cmd.AddCommand(CmdMeditationCount())

	cmd.AddCommand(CmdListScripture())
	cmd.AddCommand(CmdShowScripture())
	cmd.AddCommand(CmdListScriptureSignatureRequest())
	cmd.AddCommand(CmdShowScriptureSignatureRequest())
	cmd.AddCommand(CmdListScriptureSignatureShare())
	cmd.AddCommand(CmdShowScriptureSignatureShare())
	cmd.AddCommand(CmdListSignedScripture())
	cmd.AddCommand(CmdShowSignedScripture())
	cmd.AddCommand(CmdListSignedScriptureList())
	cmd.AddCommand(CmdShowSignedScriptureList())
	cmd.AddCommand(CmdListBlessing())
	cmd.AddCommand(CmdShowBlessing())
	cmd.AddCommand(CmdListBlessingReceipt())
	cmd.AddCommand(CmdShowBlessingReceipt())
	cmd.AddCommand(CmdListImploring())
	cmd.AddCommand(CmdShowImploring())
	cmd.AddCommand(CmdListKillConjuring())
	cmd.AddCommand(CmdShowKillConjuring())
	cmd.AddCommand(CmdListKillImploring())
	cmd.AddCommand(CmdShowKillImploring())
	cmd.AddCommand(CmdListKillMeditationSummoning())
	cmd.AddCommand(CmdShowKillMeditationSummoning())
	cmd.AddCommand(CmdListKillMagicKeySummoning())
	cmd.AddCommand(CmdShowKillMagicKeySummoning())
	cmd.AddCommand(CmdListKillScriptureSignatureRequest())
	cmd.AddCommand(CmdShowKillScriptureSignatureRequest())
	cmd.AddCommand(CmdListKillSignatureRequest())
	cmd.AddCommand(CmdShowKillSignatureRequest())
	cmd.AddCommand(CmdListEncryptedMagicKeyShare())
	cmd.AddCommand(CmdShowEncryptedMagicKeyShare())
	cmd.AddCommand(CmdListEncryptedPreSign())
	cmd.AddCommand(CmdShowEncryptedPreSign())
	// this line is used by starport scaffolding # 1

	return cmd
}
