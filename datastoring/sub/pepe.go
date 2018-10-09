package sub

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"cryptopepe.io/cryptopepe-reader/datastoring/data"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/lookups"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
)

func PepeUpdate(evCtx *event.EventContext, trig *triggers.Trigger) error {

	pepeId := trig.PepeId

	// Ignore events for pepe 0, this pepe is artificial, and burnt from the start.
	if pepeId.Uint64() == 0 {
		return nil
	}

	dataKey := convert.PepeIdToKey(pepeId)

	var err error
	var currentBlock uint64
	currentBlock, err = evCtx.Reader.GetCurrentBlock()
	if err != nil {
		return &errors.StoreError{
			Problem: "Warning! Failed to get current block number!",
			Keys: nil, Err: err}
	}

	pepeData, deletePepeData, pepeLookUpErr := lookups.PepeLookup(evCtx, trig)
	if pepeLookUpErr != nil {
		return err
	}
	saleAuctionData, deleteSaleAuctionData, saleAuctionLookupErr := lookups.SaleAuctionLookup(evCtx, trig)
	cozyAuctionData, deleteCozyAuctionData, cozyAuctionLookupErr := lookups.CozyAuctionLookup(evCtx, trig)
	// auction data may produce an error,
	//  which is most likely because there is no auction data available.
	// Disconnect/RPC-format errors ignored,
	//  but RPC spec is unclear how to differ these 2 sorts of errors.
	// TODO improve error handling

	var res data.ContentState
	if deletePepeData {
		//data doesn't exist anymore, put in a special placeholder to sign that it should be deleted.
		res = data.NewDeletableContent(currentBlock)
	} else {
		pepeData.Lcb = int64(currentBlock)

		if saleAuctionLookupErr == nil && !deleteSaleAuctionData {
			pepeData.SaleAuction = saleAuctionData
		}
		if cozyAuctionLookupErr == nil && !deleteCozyAuctionData {
			pepeData.CozyAuction = cozyAuctionData
		}

		pepeData.Update(true)
		res = pepeData
	}

	log.Printf("PepeUpdate: trig.Block: %d, current block: %d, pepe: %s\n",
		trig.Block, currentBlock, pepeId.Text(10))

	evCtx.EntityBuf.ChangeEntity(dataKey, data.ReplaceIfNewer(dataKey, res))

	return nil
}
