package reader

import (
"math/big"

"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/rpc"
)

// API is a Go wrapper around the JSON RPC API exposed by a Geth client.
type API struct {
	Client *rpc.Client
}

// request forwards an API request to the RPC server, and parses the response.
func (api *API) Request(method string, params []interface{}, dst interface{}) error {

	// Ugly hack to serialize an empty list properly
	if params == nil {
		params = []interface{}{}
	}
	if err := api.Client.Call(dst, method, params...); err != nil {
		return err
	}
	return nil
}

// BlockNumber retrieves the current head number of the blockchain.
func (api *API) BlockNumber() (uint64, error) {
	var res string
	if err := api.Request("eth_blockNumber", nil, &res); err != nil {
		return 0, err
	}
	return new(big.Int).SetBytes(common.FromHex(res)).Uint64(), nil
}
