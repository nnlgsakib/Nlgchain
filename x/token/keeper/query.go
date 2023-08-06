package keeper

import (
	"nlgcoin/x/token/types"
)

var _ types.QueryServer = Keeper{}
