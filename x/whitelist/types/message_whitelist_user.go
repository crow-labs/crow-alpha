package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWhitelistUser = "whitelist_user"

var _ sdk.Msg = &MsgWhitelistUser{}

func NewMsgWhitelistUser(creator string, address string) *MsgWhitelistUser {
	return &MsgWhitelistUser{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgWhitelistUser) Route() string {
	return RouterKey
}

func (msg *MsgWhitelistUser) Type() string {
	return TypeMsgWhitelistUser
}

func (msg *MsgWhitelistUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWhitelistUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWhitelistUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
