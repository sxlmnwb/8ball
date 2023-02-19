package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateScripture() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-scripture [index] [hash] [location] [address] [alias] [value]",
		Short: "Create a new scripture",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]
			hash := args[1]

			// Get value arguments
			argLocation := args[2]
			argAddress := args[3]
			argAlias := args[4]
			argValue := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateScripture(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				hash,
				argLocation,
				argAddress,
				argAlias,
				argValue,
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
