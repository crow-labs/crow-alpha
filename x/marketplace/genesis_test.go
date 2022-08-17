package marketplace_test

import (
	"testing"

	keepertest "github.com/crow-labs/crow/testutil/keeper"
	"github.com/crow-labs/crow/testutil/nullify"
	"github.com/crow-labs/crow/x/marketplace"
	"github.com/crow-labs/crow/x/marketplace/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketplaceKeeper(t)
	marketplace.InitGenesis(ctx, *k, genesisState)
	got := marketplace.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
