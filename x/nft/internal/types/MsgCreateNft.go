package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgCreateNft struct {
	Id      string
	Hash    string
	Creator string
	Mime    string
	Meta    string
}


func NewMsgCreateNft(id string, hash string, creator string, meta string, mime string) *MsgCreateNft {
	return &MsgCreateNft{
		Id:      id,
		Hash:    hash,
		Creator: creator,
		Mime:    mime,
		Meta:    meta,
	}
}

func (msg *MsgCreateNft) Route() string {
	return RouterKey
}

func (msg *MsgCreateNft) Type() string {
	return "CreateNft"
}

func (msg *MsgCreateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}