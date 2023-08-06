package keeper

import (
	"nlgcoin/x/multisig/types"
)

var _ types.QueryServer = Keeper{}
