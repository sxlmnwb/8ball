package cli

import (
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateVision() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vision [magic-key-id] [ec-point-x] [ec-point-y] [summoning-id]",
		Short: "Create a new vision",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMagicKeyId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argEcPointX := args[1]
			argEcPointY := args[2]
			argSummoningId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVision(clientCtx.GetFromAddress().String(), argMagicKeyId, argEcPointX, argEcPointY, argSummoningId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteVision() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-vision [id] [magic-key-id]",
		Short: "Delete a vision by id magic-key-id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			magicKeyId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVision(clientCtx.GetFromAddress().String(), magicKeyId, id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
