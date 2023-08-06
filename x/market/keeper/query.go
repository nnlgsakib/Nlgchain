package keeper

import (
	"nlgcoin/x/market/types"
)

var _ types.QueryServer = Keeper{}
