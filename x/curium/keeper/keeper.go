package keeper

import (
	"encoding/json"
	"fmt"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"time"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"

	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/core"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tendermint/tendermint/libs/log"
	"io"
	"net/http"
	"regexp"

	"github.com/bluzelle/curium/x/curium/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      *codec.Codec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		rpcPort  uint64
		accKeeper *keeper.AccountKeeper
	}
)

type MsgBroadcaster func(ctx sdk.Context, msgs []sdk.Msg, from string) chan *MsgBroadcasterResponse

type MsgBroadcasterResponse struct {
	Response *abci.ResponseDeliverTx
	Data     *[]byte
	Error    error
}

type KeyringReader struct{ keyringDir string }

func NewKeyringReader(keyringDir string) *KeyringReader {
	return &KeyringReader{
		keyringDir: keyringDir,
	}
}

func NewKeeper(cdc *codec.Codec, storeKey, memKey sdk.StoreKey, laddr string, accKeeper keeper.AccountKeeper) *Keeper {
	regex, _ := regexp.Compile(".*:")
	port, _ := math.ParseUint64(regex.ReplaceAllString(laddr, ""))
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		rpcPort:  port,
		accKeeper: &accKeeper,
	}
}

func getKeyring (keyringDir string) (cryptoKeys.Keybase, error) {
	return cryptoKeys.NewKeyring("curium", cryptoKeys.BackendTest, keyringDir, nil)
}


func getAccountAddress (keyring cryptoKeys.Keybase, from string) (sdk.AccAddress, error) {

	keyringKeys, err := keyring.Get(from)

	if err != nil {
		return nil, err
	}

	return keyringKeys.GetAddress(), nil
}

func (reader KeyringReader) GetAddress(from string) (sdk.AccAddress, error) {
	keyring, err := getKeyring(reader.keyringDir)

	if err != nil {
		return nil, err
	}

	return getAccountAddress(keyring, from)

}

func getGasPriceUbnt () (sdk.DecCoins) {
	minGasPrice := strings.Replace(viper.GetString("minimum-gas-prices"), "ubnt", "", 1)
	gasPriceDec, _ := sdk.NewDecFromStr(minGasPrice)
	return sdk.NewDecCoins(sdk.NewDecCoinFromDec("ubnt", gasPriceDec))
}

func pollForTransaction (ctx rpctypes.Context, hash []byte) (*coretypes.ResultTx, error) {

	result, err := core.Tx(&ctx, hash, false)
	if err != nil {
		time.Sleep(3 * time.Second)
		return pollForTransaction(ctx, hash)
	}

	return result, nil

}

func pollTimer (ctx rpctypes.Context, hash []byte, timeout int64) (*coretypes.ResultTx, error) {
	resultChannel := make(chan *coretypes.ResultTx, 1)
	errorChannel := make(chan error, 1)

	result, err := pollForTransaction(ctx, hash)

	errorChannel <- err
	resultChannel <- result

	select {
		case txResult := <-resultChannel:
			return txResult, nil
		case pollingError := <-errorChannel:
			return nil, pollingError
		case <-time.After(time.Duration(timeout) * time.Second):
			return nil, sdkerrors.New("curium", 1, fmt.Sprintf("Could not poll for transaction %s", string(hash)))
	}
}

func (k Keeper) NewMsgBroadcaster(keyringDir string, cdc *codec.Codec) MsgBroadcaster {
	accKeeper := k.accKeeper

	return func(ctx sdk.Context, msgs []sdk.Msg, from string) chan *MsgBroadcasterResponse {
		resp := make(chan *MsgBroadcasterResponse)

		go func() {
			returnError := func(err error) {
				resp <- &MsgBroadcasterResponse{
					Error: err,
				}
				close(resp)
			}

			kr, err := getKeyring(keyringDir)

			if err != nil {
				returnError(err)
				return
			}

			addr, err := getAccountAddress(kr, from)

			if err != nil {
				returnError(err)
				return
			}

			accnt := accKeeper.GetAccount(ctx, addr)

			// Create a new TxBuilder.
			txBuilder := auth.NewTxBuilder(
				utils.GetTxEncoder(cdc),
				accnt.GetAccountNumber(),
				accnt.GetSequence(),
				10000000,
				1,
				false,
				ctx.ChainID(),
				"memo", nil,
				getGasPriceUbnt(),
			).WithKeybase(kr)

			signedMsgs, err := txBuilder.BuildAndSign(from, clientkeys.DefaultKeyPass, msgs)

			if err != nil {
				returnError(err)
				return
			}

			if accnt == nil {
				returnError(sdkerrors.New("curium", 2, "Cannot broadcast message, accnt does not exist"))
				return
			}

			rpcCtx := rpctypes.Context{}

			broadcastResult, err := core.BroadcastTxSync(&rpcCtx, signedMsgs)

			if err != nil {
				returnError(err)
				return
			}

			result, err := pollTimer(rpcCtx, broadcastResult.Hash, 20)

			if err != nil {
				returnError(err)
				return
			}

			resp <- &MsgBroadcasterResponse{
				Response: &result.TxResult,
				Data:     &result.TxResult.Data,
			}
			close(resp)

		}()

		return resp
	}

}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

type NetInfoResult struct {
	Result NetInfo `json:"result"`
}

type NetInfo struct {
	Peers []struct {
		Ip       string `json:"remote_ip"`
		NodeInfo struct {
			Id      string `json"id"`
			Moniker string `json:"moniker"`
		} `json:"node_info"`
	} `json:"peers"`
}

type StatusResult struct {
	Result Status `json:"result"`
}
type Status struct {
	NodeInfo struct {
		Id      string `json:"id""`
		Moniker string `json:"moniker"`
	} `json:"node_info"`
}

func (k Keeper) GetStatus() (*Status, error) {
	status, err := httpGet(fmt.Sprintf("http://localhost:%d/status", k.rpcPort))
	if err != nil {
		return nil, err
	}

	var statusResult StatusResult

	json.Unmarshal(status, &statusResult)
	s := statusResult.Result
	return &s, nil
}

func (k Keeper) GetNetInfo() (*NetInfo, error) {
	info, err := httpGet(fmt.Sprintf("http://localhost:%d/net_info", k.rpcPort))
	if err != nil {
		return nil, err
	}

	var netInfoResult NetInfoResult

	json.Unmarshal(info, &netInfoResult)
	netInfo := netInfoResult.Result
	return &netInfo, nil
}

func (k Keeper) MyRemoteIp() (string, error) {
	bz, err := httpGet("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	var result map[string]string
	json.Unmarshal(bz, &result)
	return result["ip"], nil
	// TODO: NEEDED FOR LOCAL TESTING
	//return "127.0.0.1", nil
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}
