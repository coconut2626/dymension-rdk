//go:build evm
// +build evm

package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/dymensionxyz/rollapp/app"
	"github.com/dymensionxyz/rollapp/cmd/evm/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
