package sub

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"cryptopepe.io/cryptopepe-reader/datastoring/data"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
)

func UserUpdate(evCtx *event.EventContext, trig *triggers.Trigger) error {

	address := trig.Address
	dataKey := convert.AddressToUserKey(address)

	var err error
	var currentBlock uint64
	currentBlock, err = evCtx.Reader.GetCurrentBlock()
	if err != nil {
		return &errors.StoreError{
			Problem: "Warning! Failed to get current block number!",
			Keys: nil, Err: err}
	}

	caller := evCtx.Reader.GetPepeCallSession()

	deleteData := false

	var usernameBytes [32]byte
	usernameBytes, err = caller.AddressToUser(*address)
	if err != nil {
		if trig.Removed {
			deleteData = true
		} else {
			return &errors.StoreError{Problem: "Warning! Failed to retrieve username data from chain!", Keys: []string{address.Hex()}, Err: err}
		}
	}


	var res data.ContentState
	if deleteData {
		//data doesn't exist anymore, put in a special placeholder to sign that it should be deleted.
		res = data.NewDeletableContent(currentBlock)
	} else {
		// Parse data
		userData := convert.UserToUserData(usernameBytes)

		userData.Lcb = int64(currentBlock)
		res = userData
	}

	log.Printf("UserUpdate: trig.Block: %d, current block: %d, user address: %s\n",
		trig.Block, currentBlock, address.Hex())

	evCtx.EntityBuf.ChangeEntity(dataKey, data.ReplaceIfNewer(dataKey, res))

	return nil
}

