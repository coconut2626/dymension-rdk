package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dymensionxyz/rollapp/x/investment/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInvestment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "investment [tokens]",
		Short: "Query investment",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println(args)
			reqInvestor := args[0]
			reqGameBank := args[1]
			reqDenom := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryInvestmentRequest{
				Investor: reqInvestor,
				GameBank: reqGameBank,
				Denom:    reqDenom,
			}

			res, err := queryClient.Investment(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
