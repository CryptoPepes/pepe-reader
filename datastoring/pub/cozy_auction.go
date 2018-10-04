package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func CozyAuctionStartedBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetCozyAuctionEventFilterer()

	iter, err := filterer.FilterAuctionStarted(FilterOpts(fromBlock, toBlock), nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.CozyAuctionStarts <- iter.Event
	}

	iter.Close()
}


func CozyAuctionFinalizedBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetCozyAuctionEventFilterer()

	iter, err := filterer.FilterAuctionFinalized(FilterOpts(fromBlock, toBlock), nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.CozyAuctionFinalized <- iter.Event
	}

	iter.Close()

}

func CozyAuctionWonBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetCozyAuctionEventFilterer()

	iter, err := filterer.FilterAuctionWon(FilterOpts(fromBlock, toBlock), nil, nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.CozyAuctionWon <- iter.Event
	}

	iter.Close()

}

func CozyAuctionStartedWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetCozyAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionStarted(nil, eventHub.CozyAuctionStarts, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}

func CozyAuctionFinalizedWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetCozyAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionFinalized(nil, eventHub.CozyAuctionFinalized, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}

func CozyAuctionWonWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetCozyAuctionEventFilterer()

	subscription, err := filterer.WatchAuctionWon(nil, eventHub.CozyAuctionWon, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
