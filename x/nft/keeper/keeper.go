package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc         *codec.Codec
		storeKey    sdk.StoreKey
		memKey      sdk.StoreKey
		btDirectory string
		btPort      int
		homeDir     string
	}
)

func NewKeeper (
	cdc *codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	btDirectory string,
	btPort int,
	homeDir string,

) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		btDirectory: btDirectory,
		btPort:      btPort,
		homeDir:     homeDir,
	}
}
