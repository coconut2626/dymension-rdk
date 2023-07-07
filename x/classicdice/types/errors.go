package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/classicdice module sentinel errors
var (
	ErrBalanceNotEnough = sdkerrors.Register(ModuleName, 2, "balance not enough")
	ErrRngNotFound      = sdkerrors.Register(ModuleName, 3, "rng not found")
)
