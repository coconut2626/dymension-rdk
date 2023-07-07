package minipoker

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/dymensionxyz/rollapp/testutil/sample"
	minipokersimulation "github.com/dymensionxyz/rollapp/x/minipoker/simulation"
	"github.com/dymensionxyz/rollapp/x/minipoker/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = minipokersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMinipokerBetting = "op_weight_msg_minipoker_betting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMinipokerBetting int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	minipokerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&minipokerGenesis)
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

	var weightMsgMinipokerBetting int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMinipokerBetting, &weightMsgMinipokerBetting, nil,
		func(_ *rand.Rand) {
			weightMsgMinipokerBetting = defaultWeightMsgMinipokerBetting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMinipokerBetting,
		minipokersimulation.SimulateMsgMinipokerBetting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
