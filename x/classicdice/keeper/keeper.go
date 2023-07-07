package keeper

import (
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	cocotypes "github.com/dymensionxyz/rollapp/x/coco/types"
	"github.com/dymensionxyz/rollapp/x/rng"
	"github.com/tendermint/tendermint/libs/log"
	"math"
	"strconv"

	"github.com/dymensionxyz/rollapp/x/classicdice/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      storetypes.StoreKey
	memKey        storetypes.StoreKey
	paramstore    paramtypes.Subspace
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	cocoKeeper    types.CocoKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	cocoKeeper types.CocoKeeper,

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

func (k Keeper) Betting(ctx sdk.Context, msgBetting *types.MsgDiceBetting) error {
	// TODO: need to check module balance before call rng
	winChange := float64(0)
	acc := k.accountKeeper.GetAccount(ctx, msgBetting.GetSigners()[0])
	//msgSignRng := fmt.Sprintf("%s:%d:%d", msgBetting.Creator, acc.GetAccountNumber(), acc.GetSequence())
	//rngBytes := k.rng.GetRandomness(ctx, msgSignRng)
	seedRng := k.cocoKeeper.GetRng(ctx, ctx.BlockHeight())
	if seedRng == nil {
		return types.ErrRngNotFound

	}
	provablyFair := rng.NewProvablyFairRNG(rng.NewRNGConfig(msgBetting.GetSignBytes(), seedRng, acc.GetSequence()))
	rngBytes := provablyFair.Next32Bytes()
	rngUnit32 := binary.BigEndian.Uint32(rngBytes)
	winNumber := rngUnit32 % 100
	k.Logger(ctx).Debug("winNumber", "winNumber", winNumber)

	switch msgBetting.Option {
	case types.Under:
		if winNumber <= msgBetting.NumberBetting {
			winChange = float64(msgBetting.NumberBetting) / 100
		}
	case types.Over:
		if winNumber >= msgBetting.NumberBetting {
			winChange = 1 - float64(msgBetting.NumberBetting)/100
		}
	}
	payout := int64(0)
	platformFee := msgBetting.Coin.Amount.Mul(sdk.NewInt(1)).Quo(sdk.NewInt(100))
	if winChange != 0 {
		payout = int64(math.Round(0.99*(1/winChange)*10000*float64(msgBetting.Coin.Amount.Uint64())) / 10000)
	}

	k.Logger(ctx).Debug("payout", "payout", payout)

	// validate balance of game bank
	//if !k.bankKeeper.HasBalance(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName),
	//	sdk.NewCoin(msgBetting.Coin.Denom, sdk.NewInt(payout))) {
	//	return types.ErrBalanceNotEnough
	//}

	// pay token to winner and fee to game bank
	if payout > 0 {
		amountChange := sdk.NewInt(payout).Sub(platformFee)
		//err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		//	acc.GetAddress(),
		//	sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountChange)))
		//if err != nil {
		//	return err
		//}

		err := k.bankKeeper.SendCoins(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName),
			acc.GetAddress(),
			sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, amountChange)))
		if err != nil {
			return err
		}
	} else {
		amountGB := msgBetting.Coin.Amount.Sub(platformFee)
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

	if err := k.bankKeeper.SendCoins(ctx, acc.GetAddress(), k.accountKeeper.GetModuleAddress(cocotypes.PlatformAddress),
		sdk.NewCoins(sdk.NewCoin(msgBetting.Coin.Denom, platformFee))); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDiceBetting,
			sdk.NewAttribute(types.AttributeKeyCreator, msgBetting.Creator),
			sdk.NewAttribute(types.AttributeKeyWinNumber, strconv.Itoa(int(winNumber))),
			sdk.NewAttribute(types.AttributeKeyPayout, strconv.FormatInt(payout, 10)),
		),
	})

	return nil
}
