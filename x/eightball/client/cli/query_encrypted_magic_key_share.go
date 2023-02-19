package cli

import (
	"context"

	"eightball/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListEncryptedMagicKeyShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-encrypted-magic-key-share",
		Short: "list all encrypted-magic-key-share",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEncryptedMagicKeyShareRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EncryptedMagicKeyShareAll(context.Background(), params)
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

func CmdShowEncryptedMagicKeyShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-encrypted-magic-key-share [index]",
		Short: "shows a encrypted-magic-key-share",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetEncryptedMagicKeyShareRequest{
				Index: argIndex,
			}

			res, err := queryClient.EncryptedMagicKeyShare(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
