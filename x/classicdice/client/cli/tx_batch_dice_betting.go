package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dymensionxyz/rollapp/x/classicdice/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBatchDiceBetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-betting",
		Short: "Broadcast message batch diceBetting",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOption := "over"
			// convert string to uint32
			numberBetting := uint32(40)
			coin, err := sdk.ParseCoinNormalized("200urax")
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			txf, err = txf.Prepare(clientCtx)
			if err != nil {
				return err
			}

			msg := types.NewMsgDiceBetting(
				clientCtx.GetFromAddress().String(),
				argOption,
				numberBetting,
				&coin,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			startTime := time.Now()
			nextSequence := txf.Sequence()
			clientCtx.BroadcastMode = flags.BroadcastSync
			numSuccess := 0
			for i := 0; i < 100000; i++ {
				txUnsign, err := txf.BuildUnsignedTx(msg)
				if err != nil {
					return err
				}
				err = tx.Sign(txf, clientCtx.GetFromName(), txUnsign, true)
				if err != nil {
					return err
				}

				txBytes, err := clientCtx.TxConfig.TxEncoder()(txUnsign.GetTx())
				if err != nil {
					return err
				}

				// broadcast to a Tendermint node
				res, err := clientCtx.BroadcastTx(txBytes)
				if err != nil {
					return err
				}
				if res.Code == 0 {
					fmt.Println(res.TxHash)
					nextSequence = nextSequence + 1
					txf = txf.WithSequence(nextSequence)
					numSuccess++
				} else {
					time.Sleep(20 * time.Millisecond)
					accNum, accSeq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, clientCtx.GetFromAddress())
					if err != nil {
						return err
					}
					fmt.Println(fmt.Sprintf("accNum: %d, accSeq: %d", accNum, accSeq))
					txf = txf.WithSequence(accSeq)
				}
			}
			fmt.Println("numSuccess:", numSuccess)
			fmt.Println("time taken:", time.Since(startTime))

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
