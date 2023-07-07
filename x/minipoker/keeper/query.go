package keeper

import (
	"github.com/dymensionxyz/rollapp/x/minipoker/types"
)

var _ types.QueryServer = Keeper{}
