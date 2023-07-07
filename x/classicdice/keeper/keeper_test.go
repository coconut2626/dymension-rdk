package keeper

import (
	"math"
	"testing"
)

func TestPayout(t *testing.T) {
	isOver := true
	numBetting := uint32(40)
	var winChange float64
	if isOver {
		winChange = 1 - float64(numBetting)/100
	} else {
		winChange = float64(numBetting) / 100
	}
	payout := math.Round(0.99*(1/winChange*float64(10))*10000) / 10000
	t.Log(int64(payout))
}
