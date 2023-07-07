package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dymensionxyz/rollapp/x/minipoker/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMinipokerBetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minipoker-betting [coin]",
		Short: "Broadcast message minipokerBetting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCoin := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(argCoin)
			if err != nil {
				return err
			}

			msg := types.NewMsgMinipokerBetting(
				clientCtx.GetFromAddress().String(),
				&coin,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
