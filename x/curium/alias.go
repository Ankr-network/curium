package curium

import (
	"github.com/bluzelle/curium/x/curium/keeper"
	"github.com/bluzelle/curium/x/curium/types"
)

var (
	ModuleName     = types.ModuleName
	StoreKey = types.StoreKey
	MemStoreKey = types.MemStoreKey
	NewKeeper     = keeper.NewKeeper
	Broadcaster = keeper.NewBroadcaster
)
type (
	Keeper  = keeper.Keeper
)
