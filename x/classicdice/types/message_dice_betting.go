package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDiceBetting = "dice_betting"

var _ sdk.Msg = &MsgDiceBetting{}

func NewMsgDiceBetting(creator string, option string, numberBetting uint32, coin *sdk.Coin) *MsgDiceBetting {
	return &MsgDiceBetting{
		Creator:       creator,
		Option:        option,
		NumberBetting: numberBetting,
		Coin:          coin,
	}
}

func (msg *MsgDiceBetting) Route() string {
	return RouterKey
}

func (msg *MsgDiceBetting) Type() string {
	return TypeMsgDiceBetting
}

func (msg *MsgDiceBetting) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDiceBetting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDiceBetting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
