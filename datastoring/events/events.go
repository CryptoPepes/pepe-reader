package events

import (
	"cryptopepe.io/cryptopepe-reader/abi/token"
	"cryptopepe.io/cryptopepe-reader/abi/sale"
	"cryptopepe.io/cryptopepe-reader/abi/cozy"
)

type EventHub struct {

	Newborns chan *token.TokenPepeBorn
	UserNames chan *token.TokenUserNamed
	PepeNames chan *token.TokenPepeNamed
	Transfers chan *token.TokenTransfer
	SaleAuctionStarts chan *sale.SaleAuctionStarted
	SaleAuctionFinalized chan *sale.SaleAuctionFinalized
	SaleAuctionWon chan *sale.SaleAuctionWon
	CozyAuctionStarts chan *cozy.CozyAuctionStarted
	CozyAuctionFinalized chan *cozy.CozyAuctionFinalized
	CozyAuctionWon chan *cozy.CozyAuctionWon

}

func NewEventHub() *EventHub {
	hub := EventHub{
		Newborns: make(chan *token.TokenPepeBorn, 64),
		UserNames: make(chan *token.TokenUserNamed, 64),
		PepeNames: make(chan *token.TokenPepeNamed, 64),
		Transfers: make(chan *token.TokenTransfer, 64),
		SaleAuctionStarts: make(chan *sale.SaleAuctionStarted, 64),
		SaleAuctionFinalized: make(chan *sale.SaleAuctionFinalized, 64),
		SaleAuctionWon: make(chan *sale.SaleAuctionWon, 64),
		CozyAuctionStarts: make(chan *cozy.CozyAuctionStarted, 64),
		CozyAuctionFinalized: make(chan *cozy.CozyAuctionFinalized, 64),
		CozyAuctionWon: make(chan *cozy.CozyAuctionWon, 64),
	}
	return &hub
}
