package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/investment/types"
)

func (m msgServer) WithdrawGB(goCtx context.Context, msg *types.MsgWithdrawGB) (*types.MsgWithdrawGBResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get address of game bank from name
	gameBankAddress := m.GetGameBankAddress(msg.GetGameBankId())
	if gameBankAddress == nil {
		return nil, types.ErrGameBankDoesNotExist
	}

	investorAddr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return nil, err
	}

	investment, found := m.Keeper.GetInvestment(ctx, investorAddr, gameBankAddress, msg.GetAmount().GetDenom())
	if !found {
		return nil, types.ErrInvestmentDoesNotExist
	}

	// withdraw token from game bank
	amountWithdraw, err := m.Keeper.WithdrawGB(ctx, gameBankAddress, investorAddr, investment, *msg.GetAmount(), msg.GetGameBankId())
	if err != nil {
		m.Logger(ctx).Error("Withdraw GB failed", "error", err)
		return nil, err
	}

	// emit event for withdraw
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyInvestor, msg.GetCreator()),
			sdk.NewAttribute(types.AttributeKeyAmount, amountWithdraw.String()),
			sdk.NewAttribute(types.AttributeKeyShares, investment.Shares.String()),
		),
	})

	return &types.MsgWithdrawGBResponse{}, nil
}
