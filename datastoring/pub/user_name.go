package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func UserNameBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetPepeEventFilterer()
	iter, err := filterer.FilterUserNamed(FilterOpts(fromBlock, toBlock), nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.UserNames <- iter.Event
	}

	iter.Close()
}

func UserNameWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetPepeEventFilterer()

	subscription, err := filterer.WatchUserNamed(nil, eventHub.UserNames, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
