package cli

import (
	"os"

	"eightball/x/eightball/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateVerse() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-verse [from-spirit] [to-spirit] [wire-bytes] [broadcast]" +
			" [to-old] [to-old-and-new] [magic-key-id] [summoning-id]",
		Short: "Create a new verse",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFromSpirit := args[0]
			argToSpirit := strings.Split(args[1], listSeparator)
			argWireBytes := args[2]
			argBroadcast, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}
			argToOld, err := cast.ToBoolE(args[4])
			if err != nil {
				return err
			}
			argToOldAndNew, err := cast.ToBoolE(args[5])
			if err != nil {
				return err
			}
			argMagicKeyId, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}
			argSummoningId, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			dat, err := os.ReadFile(argWireBytes)
			if err != nil {
				return err
			}
			argWireBytes = string(dat)

			msg := types.NewMsgCreateVerse(clientCtx.GetFromAddress().String(), argFromSpirit, argToSpirit, argWireBytes, argBroadcast, argToOld, argToOldAndNew, argMagicKeyId, argSummoningId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
