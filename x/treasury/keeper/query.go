package keeper

import (
	"nlgcoin/x/treasury/types"
)

var _ types.QueryServer = Keeper{}
