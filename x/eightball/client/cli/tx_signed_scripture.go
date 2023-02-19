package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateSignedScripture() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signed-scripture [index] [signature] [btc-address] [magic-key-id] [sig-request-id]",
		Short: "Create a new signed-scripture",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argSignature := args[1]
			argBtcAddress := args[2]
			argMagicKeyId := args[3]

			argSigRequestId, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignedScripture(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argSignature,
				argBtcAddress,
				argMagicKeyId,
				argSigRequestId,
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
