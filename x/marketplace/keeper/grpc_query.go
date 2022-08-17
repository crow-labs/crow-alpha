package keeper

import (
	"github.com/crow-labs/crow/x/marketplace/types"
)

var _ types.QueryServer = Keeper{}
