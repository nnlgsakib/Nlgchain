package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "nlgcoin/testutil/keeper"
	"nlgcoin/x/nft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
