package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dymensionxyz/rollapp/testutil/keepers"
	"github.com/dymensionxyz/rollapp/x/investment/keeper"
	"github.com/dymensionxyz/rollapp/x/investment/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InvestmentKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
