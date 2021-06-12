package keeper

import (
	"context"
	"github.com/bluzelle/curium/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Read(goCtx context.Context, msg *types.MsgRead) (*types.MsgReadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//if !k.HasCrudValue(&ctx, msg.Uuid, msg.Key) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Key))
	//}
	val := k.GetCrudValue(&ctx, msg.Uuid, msg.Key)
	return &types.MsgReadResponse{
		Value: val.Value,
		Key:   msg.Key,
	}, nil

}
