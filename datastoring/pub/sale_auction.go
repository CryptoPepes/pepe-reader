package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func SaleAuctionStartedBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetSaleAuctionEventFilterer()

	iter, err := filterer.FilterAuctionStarted(FilterOpts(fromBlock, toBlock), nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.SaleAuctionStarts <- iter.Event
	}

	iter.Close()
}


func SaleAuctionFinalizedBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetSaleAuctionEventFilterer()

	iter, err := filterer.FilterAuctionFinalized(FilterOpts(fromBlock, toBlock), nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.SaleAuctionFinalized <- iter.Event
	}

	iter.Close()

}

func SaleAuctionWonBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetSaleAuctionEventFilterer()

	iter, err := filterer.FilterAuctionWon(FilterOpts(fromBlock, toBlock), nil, nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.SaleAuctionWon <- iter.Event
	}

	iter.Close()

}

func SaleAuctionStartedWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetSaleAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionStarted(nil, eventHub.SaleAuctionStarts, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}

func SaleAuctionFinalizedWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetSaleAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionFinalized(nil, eventHub.SaleAuctionFinalized, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}

func SaleAuctionWonWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetSaleAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionWon(nil, eventHub.SaleAuctionWon, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
