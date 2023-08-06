package keeper

import (
	"nlgcoin/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
