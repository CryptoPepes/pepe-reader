package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func PepeNameBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetPepeEventFilterer()
	iter, err := filterer.FilterPepeNamed(FilterOpts(fromBlock, toBlock), nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.PepeNames <- iter.Event
	}

	iter.Close()
}

func PepeNameWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetPepeEventFilterer()

	subscription, err := filterer.WatchPepeNamed(nil, eventHub.PepeNames, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
