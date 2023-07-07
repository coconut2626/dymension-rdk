package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/minipoker/types"
)

func (k msgServer) MinipokerBetting(goCtx context.Context, msg *types.MsgMinipokerBetting) (*types.MsgMinipokerBettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.MinipokerBetting(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgMinipokerBettingResponse{}, nil
}
