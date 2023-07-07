package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dymensionxyz/rollapp/x/classicdice/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDiceBetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dice-betting [option] [number-betting] [coin]",
		Short: "Broadcast message diceBetting",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOption := args[0]
			// convert string to uint32
			numberBetting, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}
			coin, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDiceBetting(
				clientCtx.GetFromAddress().String(),
				argOption,
				uint32(numberBetting),
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
