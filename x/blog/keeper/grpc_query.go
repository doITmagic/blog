package keeper

import (
	"github.com/doitmagic/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
