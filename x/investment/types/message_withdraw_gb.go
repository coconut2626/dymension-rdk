package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawGB = "withdraw_gb"

var _ sdk.Msg = &MsgWithdrawGB{}

func NewMsgWithdrawGB(creator string, gamebankId string, amount sdk.Coin) *MsgWithdrawGB {
	return &MsgWithdrawGB{
		Creator:    creator,
		GameBankId: gamebankId,
		Amount:     &amount,
	}
}

func (msg *MsgWithdrawGB) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawGB) Type() string {
	return TypeMsgWithdrawGB
}

func (msg *MsgWithdrawGB) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawGB) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawGB) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
