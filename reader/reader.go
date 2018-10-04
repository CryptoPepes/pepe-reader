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

func (reader *ChainReader) GetCurrentBlock() (uint64, error) {
	return reader.chainInfo.GetCurrentBlock()
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
