package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/bluzelle/curium/x/nft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc            codec.Codec
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		btClient       *torrentClient.TorrentClient
		btDirectory    string
		btPort         int
		msgBroadcaster curium.MsgBroadcaster
		homeDir        string
		keyringReader  *curium.KeyRingReader
		curiumKeeper   *curiumkeeper.Keeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	btDirectory string,
	btPort int,
	msgBroadcaster curium.MsgBroadcaster,
	homeDir string,
	keyringReader *curium.KeyRingReader,
	curiumKeeper *curiumkeeper.Keeper,
// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	btClient, err := torrentClient.NewTorrentClient(btDirectory, btPort)
	if ensureNftUserExists(keyringReader) == false {
		panic("nft user does not exist in keyring")
	}
	if err != nil {
		panic(err)
	}
	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		btClient:       btClient,
		btDirectory:    btDirectory,
		btPort:         btPort,
		msgBroadcaster: msgBroadcaster,
		homeDir:        homeDir,
		keyringReader:  keyringReader,
		curiumKeeper:   curiumKeeper,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}
