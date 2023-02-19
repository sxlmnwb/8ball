package cli

import (
	"context"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListEncryptedPreSign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-encrypted-pre-sign",
		Short: "list all encrypted-pre-sign",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEncryptedPreSignRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EncryptedPreSignAll(context.Background(), params)
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

func CmdShowEncryptedPreSign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-encrypted-pre-sign [index]",
		Short: "shows a encrypted-pre-sign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetEncryptedPreSignRequest{
				Index: argIndex,
			}

			res, err := queryClient.EncryptedPreSign(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
