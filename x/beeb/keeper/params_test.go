package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/Benzbeeb/beeb/testutil/keeper"
	"github.com/Benzbeeb/beeb/x/beeb/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BeebKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
