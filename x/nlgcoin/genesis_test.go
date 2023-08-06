package nlgcoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nlgcoin/testutil/keeper"
	"nlgcoin/testutil/nullify"
	"nlgcoin/x/nlgcoin"
	"nlgcoin/x/nlgcoin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NlgcoinKeeper(t)
	nlgcoin.InitGenesis(ctx, *k, genesisState)
	got := nlgcoin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
