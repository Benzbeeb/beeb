package beeb_test

import (
	"testing"

	keepertest "github.com/Benzbeeb/beeb/testutil/keeper"
	"github.com/Benzbeeb/beeb/testutil/nullify"
	"github.com/Benzbeeb/beeb/x/beeb"
	"github.com/Benzbeeb/beeb/x/beeb/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BeebKeeper(t)
	beeb.InitGenesis(ctx, *k, genesisState)
	got := beeb.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
