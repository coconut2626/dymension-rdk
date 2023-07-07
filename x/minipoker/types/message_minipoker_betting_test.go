package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dymensionxyz/rollapp/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMinipokerBetting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMinipokerBetting
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMinipokerBetting{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMinipokerBetting{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
