package coco

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/coco/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	msg := req.GetHash()
	rngBytes := k.GetRngFromExternal(ctx, string(msg))
	k.SetRng(ctx, rngBytes)
}
