package types

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGameBank creates a new GameBank instance
func NewGameBank(investorAddr sdk.AccAddress, gbAddr sdk.AccAddress, amount sdk.Coin, shares sdk.Dec) GameBank {
	return GameBank{
		InvestorAddress: investorAddr.String(),
		GamebankAddress: gbAddr.String(),
		Denom:           amount.Denom,
		Tokens:          amount.Amount,
		InvestorShares:  shares,
	}
}

// MustUnmarshalGameBank return the unmarshaled game bank from bytes.
// Panics if fails.
func MustUnmarshalGameBank(cdc codec.BinaryCodec, value []byte) GameBank {
	gBank, err := UnmarshalGameBank(cdc, value)
	if err != nil {
		panic(err)
	}

	return gBank
}

// UnmarshalGameBank return the delegation
func UnmarshalGameBank(cdc codec.BinaryCodec, value []byte) (gBank GameBank, err error) {
	err = cdc.Unmarshal(value, &gBank)
	return gBank, err
}

// MustMarshalGameBank marshals a game bank using the provided codec
func MustMarshalGameBank(cdc codec.BinaryCodec, gamebank *GameBank) []byte {
	return cdc.MustMarshal(gamebank)
}

// MustMarshalInvestment marshals an investment using the provided codec
func MustMarshalInvestment(cdc codec.BinaryCodec, investment *Investment) []byte {
	return cdc.MustMarshal(investment)
}

// AddTokensFromInvestor adds tokens to a game bank from an investor
func (g *GameBank) AddTokensFromInvestor(amount math.Int, refRate int64) (*GameBank, sdk.Dec) {
	// calculate the shares to issue
	var issuedShares sdk.Dec
	if g.InvestorShares.IsZero() {
		// the first investment to a game bank sets the exchange rate to one
		issuedShares = sdk.NewDec(amount.Int64())
	} else {
		// calculate amount after deduction of the reference rate
		amount = amount.Sub(amount.Mul(math.NewInt(refRate)).Quo(math.NewInt(100)))
		shares, err := g.SharesFromTokens(amount)
		if err != nil {
			panic(err)
		}

		issuedShares = shares
	}

	g.Tokens = g.Tokens.Add(amount)
	g.InvestorShares = g.InvestorShares.Add(issuedShares)

	return g, issuedShares
}

// SubTokensFromInvestor subtracts tokens from a game bank from an investor
func (g *GameBank) SubTokensFromInvestor(shares sdk.Dec) (*GameBank, math.Int, error) {
	tokens, err := g.TokensFromShares(shares)
	if err != nil {
		return nil, math.Int{}, err
	}

	if g.InvestorShares.LT(shares) {
		return nil, math.Int{}, ErrInsufficientShares
	}

	g.Tokens = g.Tokens.Sub(tokens)
	g.InvestorShares = g.InvestorShares.Sub(shares)

	return g, tokens, nil
}

// SharesFromTokens calculates the shares to issue for a given amount of tokens
func (g *GameBank) SharesFromTokens(amt math.Int) (sdk.Dec, error) {
	if g.Tokens.IsZero() {
		return sdk.NewDec(0), ErrInsufficientTokens
	}

	return g.InvestorShares.MulInt(amt).QuoInt(g.Tokens), nil
}

func (g *GameBank) TokensFromShares(shares sdk.Dec) (math.Int, error) {
	if g.InvestorShares.IsZero() {
		return math.NewInt(0), ErrInsufficientShares
	}

	return shares.Quo(g.InvestorShares).Mul(sdk.NewDecFromInt(g.Tokens)).RoundInt(), nil
}
