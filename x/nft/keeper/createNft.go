package keeper

import (
	"fmt"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/bluzelle/curium/x/nft/types"
	"os"
	"path/filepath"
	"regexp"
)



func (k Keeper) SeedFile(metainfo *metainfo.MetaInfo) error {
	err := k.BtClient.SeedFile(metainfo)
	if err != nil {
		return err
	}
	return nil
}

//func (k Keeper) broadcastPublishFile(ctx sdk.Context, id, hash string, metainfo *metainfo.MetaInfo) error{
//	metaBytes, err := bencode.EncodeBytes(metainfo)
//	if err != nil {
//		return err
//	}
//
//	addr, err := k.reader.GetAddress("nft")
//	if err != nil {
//		return err
//	}
//
//	publishMsg := types.MsgPublishFile{
//		Creator: addr.String(),
//		Id:      id,
//		Hash: hash,
//		Metainfo: metaBytes,
//	}
//
//	_, err = k.msgBroadcaster(ctx, []sdk.Msg{&publishMsg}, "nft")
//	if err != nil {
//		return err
//	}
//	return nil
//
//}

func (k Keeper) AssembleNftFile(uploadDir string, nftDir string, msg *types.MsgCreateNft) error {
	uploadRegEx, err := regexp.Compile(fmt.Sprintf("^%s-", msg.Hash))
	if err != nil {
		return err
	}

	err = filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && uploadRegEx.MatchString(info.Name()) {
			fmt.Println(path)
			if path != uploadDir {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				err = os.MkdirAll(nftDir, 0755)
				if err != nil {
					return err
				}
				f, err := os.OpenFile(nftDir+"/"+msg.Hash, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0744)
				if err != nil {
					return err
				}
				defer f.Close()
				_, err = f.Write(data)
				if err != nil {
					return err
				}
				err = os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}