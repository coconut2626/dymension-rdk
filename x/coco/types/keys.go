package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "coco"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_coco"

	// JackpotAddress defines the jackpot store key
	JackpotAddress = "jackpot_coco"

	// PlatformAddress defines the platform store key
	PlatformAddress = "coco"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	RngKey = []byte{0x21} // key for random number generator
)

// GetRngKey creates the prefix for a random number generator
func GetRngKey(blockHeight int64) []byte {
	byteArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteArray, uint64(blockHeight))
	return append(RngKey, byteArray...)
}
