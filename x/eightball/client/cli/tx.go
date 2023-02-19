package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"eightball/x/eightball/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
		Aliases:                    []string{"8ball"},
	}

	cmd.AddCommand(CmdCreateMagicKey())
	cmd.AddCommand(CmdCreateSpiritConjuringPoems())
	cmd.AddCommand(CmdFinalizeConjuring())
	cmd.AddCommand(CmdCreateVerse())
	cmd.AddCommand(CmdCreateVision())
	cmd.AddCommand(CmdDeleteVision())
	cmd.AddCommand(CmdCreateMessage())
	cmd.AddCommand(CmdCreateSignatureRequest())
	cmd.AddCommand(CmdCreateSignatureShare())
	cmd.AddCommand(CmdCreateSignedMessage())
	cmd.AddCommand(CmdCreateMeditation())
	cmd.AddCommand(CmdDeleteMeditation())
	cmd.AddCommand(CmdCreateScripture())
	cmd.AddCommand(CmdCreateScriptureSignatureRequest())
	cmd.AddCommand(CmdCreateScriptureSignatureShare())
	cmd.AddCommand(CmdCreateSignedScripture())
	cmd.AddCommand(CmdCreateSignedScriptureList())
	cmd.AddCommand(CmdCreateBlessing())
	cmd.AddCommand(CmdCreateKillConjuring())
	cmd.AddCommand(CmdCreateKillImploring())
	cmd.AddCommand(CmdCreateKillMeditationSummoning())
	cmd.AddCommand(CmdCreateKillMagicKeySummoning())
	cmd.AddCommand(CmdCreateKillScriptureSignatureRequest())
	cmd.AddCommand(CmdCreateKillSignatureRequest())
	cmd.AddCommand(CmdCreateEncryptedMagicKeyShare())
	cmd.AddCommand(CmdDeleteEncryptedMagicKeyShare())
	cmd.AddCommand(CmdCreateEncryptedPreSign())
	cmd.AddCommand(CmdDeleteEncryptedPreSign())
	// this line is used by starport scaffolding # 1

	return cmd
}
