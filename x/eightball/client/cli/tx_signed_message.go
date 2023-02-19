package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateSignedMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signed-message [body] [signature] [bitcoin-address] [message-id] [magic-key-id]",
		Short: "Create a new signed-message",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBody := args[0]
			argSignature := args[1]
			argBitcoinAddress := args[2]
			argMessageId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argSigRequestId, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			argMagicKeyId, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignedMessage(
				clientCtx.GetFromAddress().String(),
				argBody,
				argSignature,
				argBitcoinAddress,
				argMessageId,
				argSigRequestId,
				argMagicKeyId,
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
