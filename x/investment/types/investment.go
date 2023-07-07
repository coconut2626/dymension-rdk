package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewInvestment(investorAddr sdk.AccAddress, gbAddr sdk.AccAddress, denom string, shares sdk.Dec) Investment {
	return Investment{
		InvestorAddress: investorAddr.String(),
		GamebankAddress: gbAddr.String(),
		Denom:           denom,
		Shares:          shares,
	}
}

// MustUnmarshalInvestment return the unmarshaled investment from bytes.
// Panics if fails.
func MustUnmarshalInvestment(cdc codec.BinaryCodec, value []byte) Investment {
	delegation, err := UnmarshalInvestment(cdc, value)
	if err != nil {
		panic(err)
	}

	return delegation
}

// UnmarshalInvestment return the delegation
func UnmarshalInvestment(cdc codec.BinaryCodec, value []byte) (delegation Investment, err error) {
	err = cdc.Unmarshal(value, &delegation)
	return delegation, err
}
