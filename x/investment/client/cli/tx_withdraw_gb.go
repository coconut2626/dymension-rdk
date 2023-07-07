package cli

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dymensionxyz/rollapp/x/investment/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawGB() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-gb [gamebankId] [coin]",
		Short: "Broadcast message withdrawGB",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if len(args) != 2 {
				return errors.New("invalid argument count")
			}

			gamebankId := args[0]

			coin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawGB(
				clientCtx.GetFromAddress().String(),
				gamebankId,
				coin,
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
