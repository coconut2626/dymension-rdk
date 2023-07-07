package keeper

import (
	"github.com/dymensionxyz/rollapp/x/investment/types"
)

var _ types.QueryServer = Keeper{}
