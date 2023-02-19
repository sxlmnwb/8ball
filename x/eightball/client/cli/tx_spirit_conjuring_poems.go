package cli

import (
	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateSpiritConjuringPoems() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-spirit-conjuring-poems [conjurings-id] [poem]",
		Short: "Create a new spirit-conjuring-poems",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argConjuringId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argPoem := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSpiritConjuringPoems(clientCtx.GetFromAddress().String(), argConjuringId, argPoem)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
