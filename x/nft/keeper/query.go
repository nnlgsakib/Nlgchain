package keeper

import (
	"nlgcoin/x/nft/types"
)

var _ types.QueryServer = Keeper{}
