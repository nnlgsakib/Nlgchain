package keeper

import (
	"nlgcoin/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
