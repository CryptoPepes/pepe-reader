package reader

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"log"
	"cryptopepe.io/cryptopepe-reader/abi/token"
	"cryptopepe.io/cryptopepe-reader/pepe"
	"math/big"
	"cryptopepe.io/cryptopepe-reader/chainutil"
	"cryptopepe.io/cryptopepe-reader/abi/sale"
	"cryptopepe.io/cryptopepe-reader/abi/cozy"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Reader interface {
	GetCurrentBlock() (uint64, error)
	GetPepeCallSession() (*token.TokenCallerSession)
	GetSaleAuctionCallSession() (*sale.SaleCallerSession)
	GetCozyAuctionCallSession() (*cozy.CozyCallerSession)
	GetPepeEventFilterer() (*token.TokenFilterer)
	GetSaleAuctionEventFilterer() (*sale.SaleFilterer)
	GetCozyAuctionEventFilterer() (*cozy.CozyFilterer)
	GetPepeDNA(pepeId *big.Int) (pepe.PepeDNA, error)
}

type ChainReader struct {

	//For retrieving the state of the chain
	chainInfo chainutil.ChainInfo

	rawPepeToken    *token.Token
	rawSaleAuctionToken *sale.Sale
	rawCozyAuctionToken *cozy.Cozy
	pepeCallSession *token.TokenCallerSession
	saleAuctionCallSession *sale.SaleCallerSession
	cozyAuctionCallSession *cozy.CozyCallerSession

	httpCl http.Client
}

func NewChainReader(rawPepeToken *token.Token,
	rawSaleAuctionToken *sale.Sale,
	rawCozyAuctionToken *cozy.Cozy,
		chainInfo chainutil.ChainInfo) (*ChainReader) {

	reader := &ChainReader{
		chainInfo: chainInfo,
		rawPepeToken: rawPepeToken,
		rawSaleAuctionToken: rawSaleAuctionToken,
		rawCozyAuctionToken: rawCozyAuctionToken,
	}

	// Wrap the Token contract instance into a caller-session (session without auth)
	reader.pepeCallSession = &token.TokenCallerSession{
		Contract: &reader.rawPepeToken.TokenCaller,
		CallOpts: bind.CallOpts{
			Pending: false,//ignore pending data.
		},
	}

	// Wrap the Token contract instance into a caller-session (session without auth)
	reader.saleAuctionCallSession = &sale.SaleCallerSession{
		Contract: &reader.rawSaleAuctionToken.SaleCaller,
		CallOpts: bind.CallOpts{
			Pending: false,//ignore pending data.
		},
	}

	// Wrap the Token contract instance into a caller-session (session without auth)
	reader.cozyAuctionCallSession = &cozy.CozyCallerSession{
		Contract: &reader.rawCozyAuctionToken.CozyCaller,
		CallOpts: bind.CallOpts{
			Pending: false,//ignore pending data.
		},
	}

	reader.httpCl = http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	return reader
}

func (reader *ChainReader) Test() {
	//See if we can retrieve some data from the contract storage.
	name, err := reader.pepeCallSession.Name()
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)
}

type etherscanBlockNumApiResponse struct {
	BlockNumHex string `json:"result"`
}

var blockApiUrl = "https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=QV69RVEK45DAMX268Y7FC9GYQ5I87GEFFK"

func (reader *ChainReader) GetCurrentBlock() (uint64, error) {
	gethBlockNum, err := reader.chainInfo.GetCurrentBlock()
	if err != nil {
		fmt.Printf("Could not retrieve block number from geth, result: %d\n", gethBlockNum)
		fmt.Println(err)
	} else if gethBlockNum != 0 {
		return gethBlockNum, nil
	}

	req, err := http.NewRequest(http.MethodGet, blockApiUrl, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := reader.httpCl.Do(req)
	if getErr != nil {
		return 0, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return 0, readErr
	}

	resp := etherscanBlockNumApiResponse{}
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return 0, jsonErr
	}

	blockNum, parseErr := strconv.ParseInt(resp.BlockNumHex, 0, 64)
	return uint64(blockNum), parseErr

}

func (reader *ChainReader) GetPepeEventFilterer() (*token.TokenFilterer) {
	return &reader.rawPepeToken.TokenFilterer
}

func (reader *ChainReader) GetSaleAuctionEventFilterer() (*sale.SaleFilterer) {
	return &reader.rawSaleAuctionToken.SaleFilterer
}

func (reader *ChainReader) GetCozyAuctionEventFilterer() (*cozy.CozyFilterer) {
	return &reader.rawCozyAuctionToken.CozyFilterer
}

func (reader *ChainReader) GetPepeCallSession() (*token.TokenCallerSession) {
	return reader.pepeCallSession
}

func (reader *ChainReader) GetSaleAuctionCallSession() (*sale.SaleCallerSession) {
	return reader.saleAuctionCallSession
}

func (reader *ChainReader) GetCozyAuctionCallSession() (*cozy.CozyCallerSession) {
	return reader.cozyAuctionCallSession
}

func (reader *ChainReader) GetPepeDNA(pepeId *big.Int) (pepe.PepeDNA, error) {
	dnaData, err := reader.pepeCallSession.GetPepe(pepeId)
	return pepe.PepeDNA(dnaData.Genotype), err
}
