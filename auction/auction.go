package auction

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Auction struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}