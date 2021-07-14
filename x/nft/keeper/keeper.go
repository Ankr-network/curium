package keeper

import (
	"github.com/bluzelle/curium/x/curium"
	"github.com/bluzelle/curium/x/curium/keeper"
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
		msgBroadcaster curium.MsgBroadcaster
		curiumKeeper *curium.Keeper
		reader *keeper.KeyringReader
	}
)

func NewKeeper (
	cdc *codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	btDirectory string,
	btPort int,
	homeDir string,
	msgBroadcaster curium.MsgBroadcaster,
	curiumKeeper *curium.Keeper,
	reader *keeper.KeyringReader,

) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		btDirectory: btDirectory,
		btPort:      btPort,
		homeDir:     homeDir,
		msgBroadcaster: msgBroadcaster,
		curiumKeeper: curiumKeeper,
		reader: reader,
	}
}
