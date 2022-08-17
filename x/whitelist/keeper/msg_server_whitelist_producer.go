package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crow-labs/crow/x/whitelist/types"
)

func (k msgServer) WhitelistProducer(goCtx context.Context, msg *types.MsgWhitelistProducer) (*types.MsgWhitelistProducerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWhitelistProducerResponse{}, nil
}
