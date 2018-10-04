package lookups

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
	"cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/pepe"
)

func PepeLookup(evCtx *event.EventContext, trig *triggers.Trigger) (*convert.PepeData, bool, error) {

	pepeId := trig.PepeId

	caller := evCtx.Reader.GetPepeCallSession()

	pepeData, err := caller.GetPepe(pepeId)
	if err != nil {
		//if the trigger says it was removed,
		// then the error for this retrieval is to be expected
		// (but not guaranteed, a new value may already be in-place)
		if trig.Removed {
			return nil, true, nil
		} else {
			return nil, false, &errors.StoreError{
				Problem: "Warning! Failed to retrieve pepe data from chain!",
				Keys: []string{pepeId.Text(10)}, Err: err}
		}
	}

	res := pepe.Pepe(pepeData)

	// Convert data to Google Datastore format, with property keys etc.
	parsedPepe := convert.PepeToPepeData(&res, evCtx.BioGenerator)

	return parsedPepe, false, nil

}
