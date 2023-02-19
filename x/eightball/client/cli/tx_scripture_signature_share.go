package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateScriptureSignatureShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-scripture-signature-share [scripture-index] [share-data] [magic-key-id] [pub-key]",
		Short: "Create a new scripture-signature-share",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argScriptureIndex := args[0]
			argShareData := args[1]
			argMagicKeyId, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argPubKey := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateScriptureSignatureShare(clientCtx.GetFromAddress().String(), argScriptureIndex, argShareData, argMagicKeyId, argPubKey)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
