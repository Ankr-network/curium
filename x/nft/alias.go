package nft

import (
	"github.com/bluzelle/curium/x/oracle/keeper"
	"github.com/bluzelle/curium/x/oracle/types"
)

const (
	ModuleName     = types.ModuleName
	StoreKey = types.StoreKey
)

type (
	Keeper          = keeper.Keeper
)
