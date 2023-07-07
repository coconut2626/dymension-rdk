package classicdice_test

import (
	"testing"

	keepertest "github.com/dymensionxyz/rollapp/testutil/keepers"
	"github.com/dymensionxyz/rollapp/testutil/nullify"
	"github.com/dymensionxyz/rollapp/x/classicdice"
	"github.com/dymensionxyz/rollapp/x/classicdice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ClassicdiceKeeper(t)
	classicdice.InitGenesis(ctx, k, genesisState)
	got := classicdice.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
