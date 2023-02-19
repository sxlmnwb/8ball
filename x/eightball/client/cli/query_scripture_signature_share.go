package cli

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListScriptureSignatureShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-scripture-signature-share [scripture-index]",
		Short: "list all scripture-signature-share [scripture-index]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			scriptureIndexStr := args[0]

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllScriptureSignatureShareRequest{
				Pagination:     pageReq,
				ScriptureIndex: scriptureIndexStr,
			}

			res, err := queryClient.ScriptureSignatureShareAll(context.Background(), params)
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

func CmdShowScriptureSignatureShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-scripture-signature-share [scripture-index] [id]",
		Short: "shows a scripture-signature-share",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			scriptureIndexStr := args[0]

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetScriptureSignatureShareRequest{
				Id:             id,
				ScriptureIndex: scriptureIndexStr,
			}

			res, err := queryClient.ScriptureSignatureShare(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
