package cli

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListHighCouncilConjurings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-high-council-conjurings",
		Short: "list all high-council-conjurings",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHighCouncilConjuringsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HighCouncilConjuringsAll(context.Background(), params)
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

func CmdShowHighCouncilConjurings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-high-council-conjurings [id]",
		Short: "shows a high-council-conjurings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetHighCouncilConjuringsRequest{
				Id: id,
			}

			res, err := queryClient.HighCouncilConjurings(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
