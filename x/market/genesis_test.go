package market_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nlgcoin/testutil/keeper"
	"nlgcoin/testutil/nullify"
	"nlgcoin/x/market"
	"nlgcoin/x/market/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketKeeper(t)
	market.InitGenesis(ctx, *k, genesisState)
	got := market.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
