package cli

import (
	"os"
	"strconv"

	"eightball/x/eightball/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateMeditation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-meditation [from-spirit] [to-spirit] [wire-bytes] [broadcast] [to-old] [to-old-and-new] [magic-key-id] [summoning-id]",
		Short: "Create a new meditation",
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

			msg := types.NewMsgCreateMeditation(clientCtx.GetFromAddress().String(), argFromSpirit, argToSpirit, argWireBytes, argBroadcast, argToOld, argToOldAndNew, argMagicKeyId, argSummoningId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMeditation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-meditation [id]",
		Short: "Delete a meditation by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMeditation(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
