package pub

import (
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
	"github.com/ethereum/go-ethereum/event"
	"sync"
	"log"
)

type Backfiller func(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64)

type Watcher func(eventHub *events.EventHub, reader reader.Reader) event.Subscription


var backfillers = map[string]Backfiller{
	"Pepe named": PepeNameBackfill,
	"User named": UserNameBackfill,
	"Sale Auction start": SaleAuctionStartedBackfill,
	"Cozy Auction start": CozyAuctionStartedBackfill,
	"Sale Auction finalized": SaleAuctionFinalizedBackfill,
	"Cozy Auction finalized": CozyAuctionFinalizedBackfill,
	"Sale Auction won": SaleAuctionWonBackfill,
	"Cozy Auction won": CozyAuctionWonBackfill,
	"Pepe born": BornBackfill,
	"Transfer pepe": TransferBackfill,
}

func StartBackfills(eventHub *events.EventHub, reader reader.Reader, fromBlock uint64, toBlock uint64) {

	// run back-fills asynchronously, wait for completion collectively.
	var wg sync.WaitGroup

	for key, filler := range backfillers {
		wg.Add(1)
		go func(name string, fillFn Backfiller) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Backfill '%s' failed, but recovered from failure!\n", name)
				} else {
					log.Printf("Backfill '%s' finished!\n", name)
				}
			}()
			log.Printf("Backfill '%s' started!\n", name)
			fillFn(eventHub, reader, fromBlock, toBlock)
		}(key, filler)
	}

	wg.Wait()
}


var watchers = map[string]Watcher{
	"Pepe named": PepeNameWatch,
	"User named": UserNameWatch,
	"Sale Auction start": SaleAuctionStartedWatch,
	"Cozy Auction start": CozyAuctionStartedWatch,
	"Sale Auction finalized": SaleAuctionFinalizedWatch,
	"Cozy Auction finalized": CozyAuctionFinalizedWatch,
	"Sale Auction won": SaleAuctionWonWatch,
	"Cozy Auction won": CozyAuctionWonWatch,
	"Pepe born": BornWatch,
	"Transfer pepe": TransferWatch,
}

func StartWatchers(eventHub *events.EventHub, reader reader.Reader) []event.Subscription {
	watches := make([]event.Subscription, 0, len(watchers))

	for key, watcher := range watchers {

		// wrap watcher creation, and recover from potential failures.
		// (recovers RPC connection/response issues)
		watchSub := func(name string, watchFn Watcher) event.Subscription {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Failed to create watcher '%s'!\n", name)
				} else {
					log.Printf("Watch '%s' created!\n", name)
				}
			}()
			log.Printf("Watch '%s' started!\n", key)
			return watcher(eventHub, reader)
		}(key, watcher)

		watches = append(watches, watchSub)
	}

	return watches
}
