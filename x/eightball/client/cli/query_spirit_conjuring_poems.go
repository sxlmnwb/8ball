package cli

import (
	"context"
	"strconv"

	"eightball/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListSpiritConjuringPoems() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-spirit-conjuring-poems",
		Short: "list all spirit-conjuring-poems [magic-key-id]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argMagicId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSpiritConjuringPoemsRequest{
				Pagination: pageReq,
				MagicId:    argMagicId,
			}

			res, err := queryClient.SpiritConjuringPoemsAll(context.Background(), params)
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

func CmdShowSpiritConjuringPoems() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-spirit-conjuring-poems [magic-id] [id]",
		Short: "shows a spirit-conjuring-poems",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			argMagicId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetSpiritConjuringPoemsRequest{
				MagicId: argMagicId,
				Id:      id,
			}

			res, err := queryClient.SpiritConjuringPoems(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
