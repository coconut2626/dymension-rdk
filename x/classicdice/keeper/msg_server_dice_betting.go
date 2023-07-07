package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/rollapp/x/classicdice/types"
)

func (m msgServer) DiceBetting(goCtx context.Context, msg *types.MsgDiceBetting) (*types.MsgDiceBettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Betting(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgDiceBettingResponse{}, nil
}
