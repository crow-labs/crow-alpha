package keeper

import (
	"github.com/crow-labs/crow/x/whitelist/types"
)

var _ types.QueryServer = Keeper{}
