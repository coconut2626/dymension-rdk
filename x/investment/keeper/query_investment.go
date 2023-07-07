package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/investment/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Investment(goCtx context.Context, req *types.QueryInvestmentRequest) (*types.QueryInvestmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	gbAddr := k.GetGameBankAddress(req.GetGameBank())
	if gbAddr == nil {
		return nil, types.ErrGameBankDoesNotExist
	}
	gameBank, found := k.GetGameBank(ctx, gbAddr, req.GetDenom())
	if !found {
		return nil, types.ErrGameBankDoesNotExist
	}

	investorAddr, err := sdk.AccAddressFromBech32(req.GetInvestor())
	if err != nil {
		fmt.Println(fmt.Sprintf("Investment %s failed error: %v", req.GetInvestor(), err))
		return nil, err
	}
	investment, found := k.GetInvestment(ctx, investorAddr, gbAddr, req.GetDenom())
	if !found {
		return nil, types.ErrInvestmentDoesNotExist
	}

	tokens := investment.Shares.Quo(gameBank.InvestorShares).Mul(sdk.NewDecFromInt(gameBank.Tokens))

	return &types.QueryInvestmentResponse{
		Shares: investment.Shares.String(),
		Tokens: tokens.String(),
	}, nil
}
