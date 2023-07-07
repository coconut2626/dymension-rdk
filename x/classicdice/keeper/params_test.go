package keeper_test

import (
	"testing"

	testkeeper "github.com/dymensionxyz/rollapp/testutil/keepers"
	"github.com/dymensionxyz/rollapp/x/classicdice/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ClassicdiceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
