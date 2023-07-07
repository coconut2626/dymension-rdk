package keeper

import (
	"fmt"
	"github.com/chehsunliu/poker"
	cocotypes "github.com/dymensionxyz/rollapp/x/coco/types"
	"github.com/dymensionxyz/rollapp/x/rng"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/dymensionxyz/rollapp/x/minipoker/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		cocoKeeper    types.CoCoKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	cocoKeeper types.CoCoKeeper,

) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		cocoKeeper:    cocoKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MinipokerBetting(ctx sdk.Context, msgBetting *types.MsgMinipokerBetting) error {
	// TODO need validate currency bet only supported HoH coin
	acc := k.accountKeeper.GetAccount(ctx, msgBetting.GetSigners()[0])
	//msgSignRng := fmt.Sprintf("%s:%d:%d", msgBetting.Creator, acc.GetAccountNumber(), acc.GetSequence())
	//rngBytes := k.rng.GetRandomness(ctx, msgSignRng)

	seedRng := k.cocoKeeper.GetRng(ctx, ctx.BlockHeight())
	if seedRng == nil {
		return types.ErrRngNotFound

	}
	provablyFair := rng.NewProvablyFairRNG(rng.NewRNGConfig(msgBetting.GetSignBytes(), seedRng, acc.GetSequence()))
	rngBytes := provablyFair.Next32Bytes()

	// init deck
	deck := poker.NewDeckNoShuffle()
	hand := deck.DrawWithRng(5, rngBytes)
	rank := poker.RankClass(poker.Evaluate(hand))
	rankStr := poker.RankString(poker.Evaluate(hand))

	// payout
	payout := 0
	switch types.WinType(rank) {
	case types.WinTypeStraightFlush:
		// payout 50x
		payout = 50
	case types.WinTypeFourOfAKind:
		// payout 20x
		payout = 20
	case types.WinTypeFullHouse:
		// payout 10x
		payout = 10
	case types.WinTypeFlush:
		// payout 5x
		payout = 5
	case types.WinTypeStraight:
		// payout 4x
		payout = 4
	case types.WinTypeThreeOfAKind:
		// payout 3x
		payout = 3
	case types.WinTypeTwoPair:
		// payout 2x
		payout = 2
	case types.WinTypePair:
		// payout 1x
		payout = 1
	default:
		// payout 0x
		payout = 0
	}

	jackpotAmount := msgBetting.Coin.Amount.Mul(sdk.NewInt(1)).Quo(sdk.NewInt(100))
	platformFee := msgBetting.Coin.Amount.Mul(sdk.NewInt(1)).Quo(sdk.NewInt(100))

	if payout > 0 {
		amountChange := msgBetting.Coin.Amount.Mul(sdk.NewInt(int64(payout))).Sub(platformFee).Sub(jackpotAmount)
		//if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		//	acc.GetAddress(),
		//	sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountChange))); err != nil {
		//	return err
		//}

		if err := k.bankKeeper.SendCoins(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName),
			acc.GetAddress(),
			sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountChange))); err != nil {
			return err
		}
	} else {
		// subtract bet amount from user account
		amountGB := msgBetting.Coin.Amount.Sub(platformFee).Sub(jackpotAmount)
		//if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.ModuleName,
		//	sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountGB))); err != nil {
		//	return err
		//}

		if err := k.bankKeeper.SendCoins(ctx, acc.GetAddress(), k.accountKeeper.GetModuleAddress(types.ModuleName),
			sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountGB))); err != nil {
			return err
		}
	}
	//if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), cocotypes.PlatformAddress,
	//	sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, platformFee))); err != nil {
	//	return err
	//}
	//
	//if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), cocotypes.JackpotAddress,
	//	sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, jackpotAmount))); err != nil {
	//	return err
	//}

	if err := k.bankKeeper.SendCoins(ctx, acc.GetAddress(), k.accountKeeper.GetModuleAddress(cocotypes.PlatformAddress),
		sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, platformFee))); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoins(ctx, acc.GetAddress(), k.accountKeeper.GetModuleAddress(cocotypes.JackpotAddress),
		sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, jackpotAmount))); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMiniPokerBetting,
			sdk.NewAttribute(types.AttributeKeyCreator, msgBetting.Creator),
			sdk.NewAttribute(types.AttributeKeyWinType, rankStr),
			sdk.NewAttribute(types.AttributeKeyJackpot, jackpotAmount.String()),
			sdk.NewAttribute(types.AttributeKeyPayout, strconv.Itoa(payout)),
		),
	})

	return nil
}
