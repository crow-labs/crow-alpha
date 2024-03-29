package whitelist

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/crow-labs/crow/testutil/sample"
	whitelistsimulation "github.com/crow-labs/crow/x/whitelist/simulation"
	"github.com/crow-labs/crow/x/whitelist/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = whitelistsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgWhitelistUser = "op_weight_msg_whitelist_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWhitelistUser int = 100

	opWeightMsgWhitelistProducer = "op_weight_msg_whitelist_producer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWhitelistProducer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	whitelistGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&whitelistGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgWhitelistUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWhitelistUser, &weightMsgWhitelistUser, nil,
		func(_ *rand.Rand) {
			weightMsgWhitelistUser = defaultWeightMsgWhitelistUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWhitelistUser,
		whitelistsimulation.SimulateMsgWhitelistUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWhitelistProducer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWhitelistProducer, &weightMsgWhitelistProducer, nil,
		func(_ *rand.Rand) {
			weightMsgWhitelistProducer = defaultWeightMsgWhitelistProducer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWhitelistProducer,
		whitelistsimulation.SimulateMsgWhitelistProducer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
