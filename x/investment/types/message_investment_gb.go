package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInvestmentGB = "investment_gb"

var _ sdk.Msg = &MsgInvestmentGB{}

func NewMsgInvestmentGB(creator string, gamebankId string, amount sdk.Coin) *MsgInvestmentGB {
	return &MsgInvestmentGB{
		Creator:    creator,
		GameBankId: gamebankId,
		Coin:       &amount,
	}
}

func (msg *MsgInvestmentGB) Route() string {
	return RouterKey
}

func (msg *MsgInvestmentGB) Type() string {
	return TypeMsgInvestmentGB
}

func (msg *MsgInvestmentGB) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInvestmentGB) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInvestmentGB) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
