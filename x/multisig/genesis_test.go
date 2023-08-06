package multisig_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nlgcoin/testutil/keeper"
	"nlgcoin/testutil/nullify"
	"nlgcoin/x/multisig"
	"nlgcoin/x/multisig/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MultisigKeeper(t)
	multisig.InitGenesis(ctx, *k, genesisState)
	got := multisig.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
