package triggers

import (
	"math/big"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
	"sync"
	"github.com/ethereum/go-ethereum/common"
)

type TriggerListener chan *Trigger

type TriggerBroadcast struct {
	sync.Mutex
	listeners []TriggerListener
}

func NewTriggerBroadcast() *TriggerBroadcast {
	br := new(TriggerBroadcast)
	br.listeners = make([]TriggerListener, 0)
	return br
}

func (broadcast *TriggerBroadcast) RegisterListener(listener TriggerListener) {
	//lock the broadcast while registering a new listener
	broadcast.Lock()
	defer broadcast.Unlock()

	//add the listener
	broadcast.listeners = append(broadcast.listeners, listener)
}

func (broadcast *TriggerBroadcast) Broadcast(trig *Trigger) {
	//lock broadcast, no changes in listeners during fan out
	broadcast.Lock()
	defer broadcast.Unlock()

	//fan out the message to all listeners
	for _, li := range broadcast.listeners {
		li <- trig
	}
}


func (broadcast *TriggerBroadcast) Close() {
	broadcast.Lock()
	defer broadcast.Unlock()

	//close all channels
	for _, li := range broadcast.listeners {
		close(li)
	}

	//remove all listeners,
	broadcast.listeners = nil
}

type Trigger struct {

	// May be nil
	PepeId *big.Int

	// May be nil
	Address *common.Address

	Block uint64

	//True when the trigger is fired to undo itself. (When TX logs are rolled back)
	Removed bool
}

func NewTrigger(pepeId *big.Int, address *common.Address, block uint64) *Trigger {
	return &Trigger{
		PepeId: pepeId,
		Address: address,
		Block: block,
	}
}

type TriggerHub struct {
	Pepe *TriggerBroadcast
	User *TriggerBroadcast
	quit chan int
}


func NewTriggerHub() *TriggerHub {
	hub := &TriggerHub{
		Pepe: NewTriggerBroadcast(),
		User: NewTriggerBroadcast(),
	}
	return hub
}

/**
Map events to triggers.
Note that some events can trigger multiple things, and these triggers can have multiple listeners.
 */
func (triggerHub *TriggerHub) TriggerMap(eventHub *events.EventHub) {

	//TODO: add more events, some triggers may need special transforms.

	// Using generics would be way cleaner, but Go makes this design type hard.
	// At least we can move the triggering logic to this function to remove code replication a bit.
	auctionHandler := func(pepe *big.Int, block uint64) {
		trig := NewTrigger(pepe, nil, block)
		// update price data, which is a part of the pepe data, so update the whole entity.
		triggerHub.Pepe.Broadcast(trig)
	}

	// process event->trigger mapping in background
	go func() {
		for {
			select {
			case newborn := <-eventHub.Newborns:
				trig := NewTrigger(newborn.PepeId, nil, newborn.Raw.BlockNumber)
				triggerHub.Pepe.Broadcast(trig)
				// Also update the mother and father, since their cooldown changes.
				// If there is no mother/father, or if it's 0, then don't update the father/mother.
				if newborn.Mother != nil && newborn.Mother.Uint64() != 0 {
					motherTrig := NewTrigger(newborn.Mother, nil, newborn.Raw.BlockNumber)
					triggerHub.Pepe.Broadcast(motherTrig)
				}
				if newborn.Father != nil && newborn.Father.Uint64() != 0 {
					fatherTrig := NewTrigger(newborn.Father, nil, newborn.Raw.BlockNumber)
					triggerHub.Pepe.Broadcast(fatherTrig)
				}
			case ev := <-eventHub.PepeNames:
				trig := NewTrigger(ev.PepeId, nil, ev.Raw.BlockNumber)
				triggerHub.Pepe.Broadcast(trig)
			case ev := <-eventHub.SaleAuctionStarts:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case ev := <-eventHub.CozyAuctionStarts:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case ev := <-eventHub.SaleAuctionFinalized:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case ev := <-eventHub.CozyAuctionFinalized:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case ev := <-eventHub.SaleAuctionWon:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case ev := <-eventHub.CozyAuctionWon:
				auctionHandler(ev.Pepe, ev.Raw.BlockNumber)
			case transfer := <-eventHub.Transfers:
				trig := NewTrigger(transfer.TokenId, nil, transfer.Raw.BlockNumber)
				triggerHub.Pepe.Broadcast(trig)
			case ev := <-eventHub.UserNames:
				trig := NewTrigger(nil, &ev.User, ev.Raw.BlockNumber)
				triggerHub.User.Broadcast(trig)
			case <-triggerHub.quit:
				return
			}
		}
	}()
}

func (hub *TriggerHub) Stop() {
	hub.quit <- 0
	hub.Pepe.Close()
}

