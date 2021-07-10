package nft

import (
	keeper2 "github.com/bluzelle/curium/x/nft/keeper"
	types2 "github.com/bluzelle/curium/x/nft/types"
)

var (
	ModuleName     = types2.ModuleName
	StoreKey = types2.StoreKey
	MemStoreKey = types2.MemStoreKey
	NewKeeper     = keeper2.NewKeeper

)
type (
	Keeper  = keeper2.Keeper
)
