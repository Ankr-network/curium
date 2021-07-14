// Copyright (C) 2020 Bluzelle
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License, version 3,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package nft

import (
	"encoding/json"
	"fmt"
	"github.com/bluzelle/curium/x/nft/keeper"
	"github.com/bluzelle/curium/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"os"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgCreateNft:
			return handleMsgCreateNft(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nft msg type: %v", msg.Type()))
		}
	}
}

func handleMsgCreateNft(goCtx sdk.Context, k keeper.Keeper, msg *types.MsgCreateNft) (*sdk.Result, error) {

	k.AppendNft(
		goCtx,
		msg.Creator,
		msg.Id,
		msg.Hash,
		msg.Vendor,
		msg.UserId,
		msg.Meta,
		msg.Mime,
	)
	err := k.AssembleNftFile(k.HomeDir+"/nft-upload", k.HomeDir+"/nft", msg)
	if err != nil {
		return nil, sdkerrors.New("nft", 2, fmt.Sprintf("unable to move nft files: %s", msg.Hash))
	}



	if _, err := os.Stat(k.HomeDir+"/nft/" + msg.Hash); err == nil {
		metainfo, err := k.BtClient.TorrentFromFile(msg.Hash)
		if err != nil {
			return nil, sdkerrors.New("nft", 2, fmt.Sprintf("unable to create torrent for file", msg.Hash))
		}
		err = k.SeedFile(metainfo)
		if err != nil {
			return nil, sdkerrors.New("nft", 2, fmt.Sprintf("unable to seed file: %s", msg.Hash))
		}

		//go func() {
		//	err = k.broadcastPublishFile(goCtx, msg.Id, msg.Hash, metainfo)
		//	if err != nil {
		//		k.Logger(goCtx).Error("error broadcasting publish nft file", "err", err.Error())
		//	}
		//}()
	}

	if err != nil {
		return nil, sdkerrors.New("nft", 2, fmt.Sprintf("unable to create torrent:  %s", msg.Hash))
	}



	createResp, err := json.Marshal(types.MsgCreateNftResponse{
		Id: msg.Id,
	})

	if err != nil {
		return nil, sdkerrors.New("nft", 2, "Failed to marshal MsgCreateNftResponse")
	}

	return &sdk.Result{Data: createResp}, nil
}