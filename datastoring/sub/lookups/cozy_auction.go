package lookups

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
	"cryptopepe.io/cryptopepe-reader/auction"
)

// Returns: data, if data was deleted, possible errors
func CozyAuctionLookup(evCtx *event.EventContext, trig *triggers.Trigger) (*convert.AuctionData, bool, error) {

	pepeId := trig.PepeId

	caller := evCtx.Reader.GetCozyAuctionCallSession()

	auctionData, err := caller.Auctions(pepeId)
	if err != nil {
		if trig.Removed {
			return nil, true, nil
		} else {
			return nil, false, &errors.StoreError{Problem: "Warning! Failed to retrieve auction data from chain!", Keys: []string{pepeId.Text(10)}, Err: err}
		}
	}

	res := auction.Auction(auctionData)

	if res.AuctionBegin == 0 {
		// This is only 0 in non-existing auctions (0s-return from RPC)
		return nil, true, nil
	}

	// Convert data to Google Datastore format, with property keys etc.
	parsedAuction := convert.AuctionToAuctionData(&res)

	return parsedAuction, false, nil
}
