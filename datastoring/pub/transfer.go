package pub

import (
	"github.com/ethereum/go-ethereum/event"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
)

func TransferBackfill(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	filterer := reader.GetPepeEventFilterer()
	iter, err := filterer.FilterTransfer(FilterOpts(fromBlock, toBlock), nil, nil, nil)
	if err != nil {
		panic(err)
	}

	for iter.Next() {
		eventHub.Transfers <- iter.Event
	}

	iter.Close()
}

func TransferWatch(eventHub *events.EventHub, reader reader.Reader) event.Subscription {

	filterer := reader.GetPepeEventFilterer()

	subscription, err := filterer.WatchTransfer(nil, eventHub.Transfers, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	return subscription
}
