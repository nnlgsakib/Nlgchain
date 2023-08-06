package keeper

import (
	"nlgcoin/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
