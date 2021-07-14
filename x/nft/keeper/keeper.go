package keeper

import (
	"fmt"
	"github.com/bluzelle/curium/x/curium"
	"github.com/bluzelle/curium/x/curium/keeper"
	"github.com/bluzelle/curium/x/nft/types"
	"github.com/bluzelle/curium/x/torrentClient"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
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
		btClient       *torrentClient.TorrentClient
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
	btClient, err := torrentClient.NewTorrentClient(btDirectory, btPort)
	if ensureNftUserExists(reader) == false {
		panic("nft user does not exist in keyring")
	}
	if err != nil {
		panic(err)
	}
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
		btClient: btClient,
	}
}

func ensureNftUserExists(reader *curium.KeyringReader) bool {
	_, err := reader.GetAddress("nft")
	return err == nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetPeerStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerKey))
}

func (k Keeper) CheckIsNftAdmin(address string) error {
	nftAdmin, err := k.reader.GetAddress("nft")
	if err != nil {
		return sdkerrors.New("nft", 1, fmt.Sprintf("peer request with invalid admin: %s", err.Error()))
	}
	if address != nftAdmin.String() {
		return sdkerrors.New("nft", 1, fmt.Sprintf("peer request with invalid admin: %s", address))
	}
	return nil
}

func (k Keeper) GetMyNodeId(ctx sdk.Context) string {
	status, err := k.curiumKeeper.GetStatus()
	if err != nil {
		logger := k.Logger(ctx)
		logger.Error("unable to get node id", err)
	}
	return status.NodeInfo.Id
}

//func (k Keeper) BroadcastRegisterBtPeer(ctx sdk.Context) {
//	nodeId := k.GetMyNodeId(ctx)
//
//	myIp, err := k.curiumKeeper.MyRemoteIp()
//	if err != nil || myIp == "" {
//		k.Logger(ctx).Error("unable to get my ip", err)
//	}
//
//	creator, err := k.keyringReader.GetAddress("nft")
//	if err != nil {
//		k.Logger(ctx).Error("unable to get address of nft account")
//	}
//	msg := types.MsgRegisterPeer{
//		Creator: creator.String(),
//		Id:      nodeId,
//		Address: myIp,
//		Port:    uint64(k.btPort),
//	}
//
//	result :=  <- k.msgBroadcaster(ctx, []sdk.Msg{&msg}, "nft")
//
//	if result.Error != nil {
//		k.Logger(ctx).Error("unable to broadcast register peer message", err)
//	}
//}
