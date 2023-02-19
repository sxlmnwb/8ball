package keeper

import (
	"eightball/x/eightball/types"
)

var _ types.QueryServer = Keeper{}
