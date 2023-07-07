package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/investment module sentinel errors
var (
	ErrInsufficientShares     = errors.Register(ModuleName, 2, "insufficient investor shares")
	ErrGameBankDoesNotExist   = errors.Register(ModuleName, 3, "gamebank does not exist")
	ErrInvestmentDoesNotExist = errors.Register(ModuleName, 4, "investment does not exist")
	ErrInsufficientTokens     = errors.Register(ModuleName, 5, "insufficient tokens")
)
