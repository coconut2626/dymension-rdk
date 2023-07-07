package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "investment"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_investment"
)

var (
	InvestorKey = []byte{0x21} // key for a investment
	GameBankKey = []byte{0x31} // key for a game bank
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetGameBankKey(gameBankAddr sdk.AccAddress, denom string) []byte {
	keyBytes := append(GameBankKey, []byte(denom)...)
	return append(keyBytes, address.MustLengthPrefix(gameBankAddr)...)
}

func GetInvestorKey(investorAddr sdk.AccAddress, gameBankAddr sdk.AccAddress, denom string) []byte {
	return append(append(GetDelegationsKey(investorAddr), []byte(denom)...), address.MustLengthPrefix(gameBankAddr)...)
}

// GetDelegationsKey creates the prefix for a investor for all game banks
func GetDelegationsKey(investorAddr sdk.AccAddress) []byte {
	return append(InvestorKey, address.MustLengthPrefix(investorAddr)...)
}
