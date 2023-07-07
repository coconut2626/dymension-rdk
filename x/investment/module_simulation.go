package investment

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/dymensionxyz/rollapp/testutil/sample"
	investmentsimulation "github.com/dymensionxyz/rollapp/x/investment/simulation"
	"github.com/dymensionxyz/rollapp/x/investment/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = investmentsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgInvestmentGB = "op_weight_msg_investment_gb"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInvestmentGB int = 100

	opWeightMsgWithdrawGB = "op_weight_msg_withdraw_gb"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawGB int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	investmentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&investmentGenesis)
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

	var weightMsgInvestmentGB int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInvestmentGB, &weightMsgInvestmentGB, nil,
		func(_ *rand.Rand) {
			weightMsgInvestmentGB = defaultWeightMsgInvestmentGB
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInvestmentGB,
		investmentsimulation.SimulateMsgInvestmentGB(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawGB int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawGB, &weightMsgWithdrawGB, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawGB = defaultWeightMsgWithdrawGB
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawGB,
		investmentsimulation.SimulateMsgWithdrawGB(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
