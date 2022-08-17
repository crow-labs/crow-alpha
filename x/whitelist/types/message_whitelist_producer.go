package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWhitelistProducer = "whitelist_producer"

var _ sdk.Msg = &MsgWhitelistProducer{}

func NewMsgWhitelistProducer(creator string, address string) *MsgWhitelistProducer {
	return &MsgWhitelistProducer{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgWhitelistProducer) Route() string {
	return RouterKey
}

func (msg *MsgWhitelistProducer) Type() string {
	return TypeMsgWhitelistProducer
}

func (msg *MsgWhitelistProducer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWhitelistProducer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWhitelistProducer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
