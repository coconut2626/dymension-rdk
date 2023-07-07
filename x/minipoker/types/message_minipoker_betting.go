package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMinipokerBetting = "minipoker_betting"

var _ sdk.Msg = &MsgMinipokerBetting{}

func NewMsgMinipokerBetting(creator string, coin *sdk.Coin) *MsgMinipokerBetting {
	return &MsgMinipokerBetting{
		Creator: creator,
		Coin:    coin,
	}
}

func (msg *MsgMinipokerBetting) Route() string {
	return RouterKey
}

func (msg *MsgMinipokerBetting) Type() string {
	return TypeMsgMinipokerBetting
}

func (msg *MsgMinipokerBetting) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMinipokerBetting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMinipokerBetting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
