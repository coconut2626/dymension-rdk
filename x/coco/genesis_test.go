package coco_test

import (
	"testing"

	keepertest "github.com/dymensionxyz/rollapp/testutil/keepers"
	"github.com/dymensionxyz/rollapp/testutil/nullify"
	"github.com/dymensionxyz/rollapp/x/coco"
	"github.com/dymensionxyz/rollapp/x/coco/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CocoKeeper(t)
	coco.InitGenesis(ctx, k, genesisState)
	got := coco.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
