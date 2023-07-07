package keeper

import (
	"cosmossdk.io/math"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/dymensionxyz/rollapp/x/investment/types"
)

const RefRate = 2 // 2% per investment

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,

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
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InvestmentGB(ctx sdk.Context, gbAddress sdk.AccAddress, investor sdk.AccAddress,
	amount sdk.Coin, gameBankId string) (sdk.Dec, error) {
	// Get or create the delegation object
	investment, found := k.GetInvestment(ctx, investor, gbAddress, amount.GetDenom())
	if !found {
		investment = types.NewInvestment(investor, gbAddress, amount.GetDenom(), sdk.NewDec(0))
	}

	// Get or create the game bank object
	gameBank, found := k.GetGameBank(ctx, gbAddress, amount.GetDenom())
	if !found {
		gameBank = types.NewGameBank(investor, gbAddress, sdk.NewCoin(amount.Denom, sdk.NewInt(0)), sdk.NewDec(0))
	}

	// send tokens from investor to game bank
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, investor, gameBankId, sdk.NewCoins(amount))
	if err != nil {
		return sdk.Dec{}, err
	}

	gbUpdate, newShare := gameBank.AddTokensFromInvestor(amount.Amount, RefRate)

	// Update the game bank object
	k.SetGameBank(ctx, *gbUpdate, gbAddress)

	// Update the investment object
	investment.Shares = investment.Shares.Add(newShare)
	k.SetInvestment(ctx, investment, investor, gbAddress)

	return newShare, nil
}

// WithdrawGB withdraws a specific investment of investor.
func (k Keeper) WithdrawGB(ctx sdk.Context, gameBankAddress sdk.AccAddress,
	investorAddr sdk.AccAddress, investment types.Investment, amount sdk.Coin, gameBankId string) (math.Int, error) {
	gameBank, found := k.GetGameBank(ctx, gameBankAddress, amount.Denom)
	if !found {
		return math.NewInt(0), types.ErrGameBankDoesNotExist
	}

	// Calculate share that withdraw from investor
	share := gameBank.InvestorShares.MulInt(amount.Amount).Quo(sdk.NewDecFromInt(gameBank.Tokens))
	if share.GT(investment.Shares) {
		share = investment.Shares
	}

	gbUpdate, amountWithdraw, err := gameBank.SubTokensFromInvestor(share)
	if err != nil {
		return math.NewInt(0), err
	}

	// Send token from game bank to investor
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, gameBankId,
		investorAddr, sdk.NewCoins(sdk.NewCoin(investment.Denom, amountWithdraw)))
	if err != nil {
		k.Logger(ctx).Error("WithdrawGB", "err", err)
		return math.NewInt(0), err
	}

	// Update the game bank object
	k.SetGameBank(ctx, *gbUpdate, gameBankAddress)

	// Update the investment object
	investment.Shares = investment.Shares.Sub(share)
	k.SetInvestment(ctx, investment, investorAddr, gameBankAddress)

	return amountWithdraw, nil
}

// GetGameBankAddress returns the game bank module account address.
func (k Keeper) GetGameBankAddress(moduleName string) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(moduleName)
}

// GetInvestment returns a specific investment of investor.
func (k Keeper) GetInvestment(ctx sdk.Context, investorAddr sdk.AccAddress,
	gameBankAddr sdk.AccAddress, denom string) (investment types.Investment, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetInvestorKey(investorAddr, gameBankAddr, denom)

	value := store.Get(key)
	if value == nil {
		return investment, false
	}

	investment = types.MustUnmarshalInvestment(k.cdc, value)

	return investment, true
}

func (k Keeper) GetGameBank(ctx sdk.Context,
	gamebankAddr sdk.AccAddress,
	denom string) (gamebank types.GameBank, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetGameBankKey(gamebankAddr, denom)

	value := store.Get(key)
	if value == nil {
		return gamebank, false
	}
	gamebank = types.MustUnmarshalGameBank(k.cdc, value)

	return gamebank, true
}

// SetGameBank sets a specific game bank of investor.
func (k Keeper) SetGameBank(ctx sdk.Context, gamebank types.GameBank, gamebankAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalGameBank(k.cdc, &gamebank)
	store.Set(types.GetGameBankKey(gamebankAddr, gamebank.Denom), bz)
}

// SetInvestment sets a specific investment of investor.
func (k Keeper) SetInvestment(ctx sdk.Context, investment types.Investment,
	investor sdk.AccAddress, gbAddress sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalInvestment(k.cdc, &investment)
	store.Set(types.GetInvestorKey(investor, gbAddress, investment.Denom), b)
}
