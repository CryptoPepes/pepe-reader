package reader

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type ClientChainInfo struct {
	client *rpc.Client
	Api *API
}

func NewClientChainInfo(rpcClient *rpc.Client) *ClientChainInfo {
	api := &API{Client: rpcClient}
	info := &ClientChainInfo{
		client: rpcClient,
		Api: api,
	}
	return info
}

func (clientInfo *ClientChainInfo) GetCurrentBlock() (uint64, error) {
	return clientInfo.Api.BlockNumber()
}
