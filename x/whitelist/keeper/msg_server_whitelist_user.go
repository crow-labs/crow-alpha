package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crow-labs/crow/x/whitelist/types"
)

func (k msgServer) WhitelistUser(goCtx context.Context, msg *types.MsgWhitelistUser) (*types.MsgWhitelistUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWhitelistUserResponse{}, nil
}
