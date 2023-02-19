package cli

import (
	"eightball/x/eightball/types"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateEncryptedPreSign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-encrypted-pre-sign [magic-key-id] [data-file]",
		Short: "Create a new encrypted-pre-sign",
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

			msg := types.NewMsgCreateEncryptedPreSign(
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

func CmdDeleteEncryptedPreSign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-encrypted-pre-sign [index]",
		Short: "Delete a encrypted-pre-sign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteEncryptedPreSign(
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
