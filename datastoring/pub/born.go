package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func BornBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetPepeEventFilterer()
	iter, err := filterer.FilterPepeBorn(FilterOpts(fromBlock, toBlock), nil, nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.Newborns <- iter.Event
	}

	iter.Close()
}

func BornWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetPepeEventFilterer()

	subscription, err := filterer.WatchPepeBorn(nil, eventHub.Newborns, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
