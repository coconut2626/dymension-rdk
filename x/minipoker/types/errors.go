package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/minipoker module sentinel errors
var (
	ErrRngNotFound = sdkerrors.Register(ModuleName, 3, "rng not found")
)
