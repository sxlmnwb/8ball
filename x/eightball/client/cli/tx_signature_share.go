package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateSignatureShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signature-share [message-id] [share-data] [magic-key-id] [public-key-string]",
		Short: "Create a new signature-share",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMessageId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argShareData := args[1]
			argMagicKeyId, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argPubKeyString := args[3]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignatureShare(clientCtx.GetFromAddress().String(), argMessageId, argShareData, argMagicKeyId, argPubKeyString)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
