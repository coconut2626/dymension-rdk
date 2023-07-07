package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/investment/types"
)

func (m msgServer) InvestmentGB(goCtx context.Context, msg *types.MsgInvestmentGB) (*types.MsgInvestmentGBResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get address of game bank from name
	gameBankAddress := m.GetGameBankAddress(msg.GetGameBankId())
	if gameBankAddress == nil {
		return nil, types.ErrGameBankDoesNotExist
	}

	// TODO: check game bank status and policy

	investorAddr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		m.Logger(ctx).Error("Investment GB failed", "error", err)
		return nil, err
	}

	newShares, err := m.Keeper.InvestmentGB(ctx, gameBankAddress, investorAddr, *msg.GetCoin(),
		msg.GetGameBankId())
	if err != nil {
		m.Logger(ctx).Error("Investment GB failed", "error", err)
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeInvestment,
			sdk.NewAttribute(types.AttributeKeyInvestor, msg.GetCreator()),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.GetCoin().String()),
			sdk.NewAttribute(types.AttributeKeyShares, newShares.String()),
		),
	})

	return &types.MsgInvestmentGBResponse{}, nil
}
