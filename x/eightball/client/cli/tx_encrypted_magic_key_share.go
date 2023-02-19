package cli

import (
	"eightball/x/eightball/types"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateEncryptedMagicKeyShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-encrypted-magic-key-share [magic-key-id] [data-file]",
		Short: "Create a new encrypted-magic-key-share",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			magicKeyStr := args[0]

			// Get value arguments
			argDataFile := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			dat, err := os.ReadFile(argDataFile)
			if err != nil {
				return err
			}
			argData := string(dat)

			msg := types.NewMsgCreateEncryptedMagicKeyShare(
				clientCtx.GetFromAddress().String(),
				magicKeyStr,
				argData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteEncryptedMagicKeyShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-encrypted-magic-key-share [index]",
		Short: "Delete a encrypted-magic-key-share",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteEncryptedMagicKeyShare(
				clientCtx.GetFromAddress().String(),
				indexIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
