package types

type WinType int32

const (
	WinTypeStraightFlush WinType = 1
	WinTypeFourOfAKind           = 2
	WinTypeFullHouse             = 3
	WinTypeFlush                 = 4
	WinTypeStraight              = 5
	WinTypeThreeOfAKind          = 6
	WinTypeTwoPair               = 7
	WinTypePair                  = 8
	WinTypeHighCard              = 9
)

const (
	EventTypeMiniPokerBetting = "minipoker_betting"

	AttributeKeyCreator = "creator"
	AttributeKeyWinType = "win_type"
	AttributeKeyJackpot = "jackpot"
	AttributeKeyPayout  = "payout"
)
