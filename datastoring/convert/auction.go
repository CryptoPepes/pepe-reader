package convert

import (
	auctionSpec "cryptopepe.io/cryptopepe-reader/auction"
	"strings"
	"fmt"
)

func AuctionToAuctionData(auction *auctionSpec.Auction) *AuctionData {
	data := new(AuctionData)
	// Converting to signed int is fine, as the sign-bit is never used anyway.
	data.BeginTime = int64(auction.AuctionBegin)
	data.EndTime = int64(auction.AuctionEnd)

	// Decimal, padded string. Padding is for alphabetic sorting to work the same as numeric.
	// 39 decimals: log(2^128) = ~ 38.53, round up
	data.BeginPrice = fmt.Sprintf("%039s", auction.BeginPrice.String())
	data.EndPrice = fmt.Sprintf("%039s", auction.EndPrice.String())

	data.Seller = strings.ToLower(auction.Seller.Hex())

	return data
}

type AuctionData struct {

	// Full price, to prevent rounding errors. Decimal.
	// Padded with 0s to 39 chars (log(2^128) = 38.5....).
	BeginPrice string `datastore:"begin_price"`
	EndPrice   string `datastore:"end_price"`

	BeginTime int64  `datastore:"begin_time"`
	EndTime   int64  `datastore:"end_time"`
	Seller    string `datastore:"seller"`

}

func (auction *AuctionData) IsExpired(time int64) bool {
	return auction.EndTime < time
}
