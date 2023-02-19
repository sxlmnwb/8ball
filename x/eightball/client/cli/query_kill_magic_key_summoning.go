package cli

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListKillMagicKeySummoning() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-kill-magic-key-summoning",
		Short: "list all kill-magic-key-summoning",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllKillMagicKeySummoningRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.KillMagicKeySummoningAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowKillMagicKeySummoning() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-kill-magic-key-summoning [id]",
		Short: "shows a kill-magic-key-summoning",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetKillMagicKeySummoningRequest{
				Id: id,
			}

			res, err := queryClient.KillMagicKeySummoning(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
