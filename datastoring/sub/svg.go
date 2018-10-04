package sub

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/pepe"
	"cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"cryptopepe.io/cryptopepe-reader/datastoring/data"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
)

func SvgUpdate(evCtx *event.EventContext, trig *triggers.Trigger) error {

	pepeId := trig.PepeId
	dataKey := convert.PepeIdToSVGKey(pepeId)

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

	var pepeData pepe.Pepe
	pepeData, err = caller.GetPepe(pepeId)
	if err != nil {
		if trig.Removed {
			deleteData = true
		} else {
			return &errors.StoreError{Problem: "Warning! Failed to retrieve pepe data from chain!", Keys: []string{pepeId.Text(10)}, Err: err}
		}
	}


	var res data.ContentState
	if deleteData {
		//data doesn't exist anymore, put in a special placeholder to sign that it should be deleted.
		res = data.NewDeletableContent(currentBlock)
	} else {
		// Parse data
		parsedPepe := convert.PepeToPepeData(&pepeData, evCtx.BioGenerator)

		svgData, err := convert.PepeToSVG(&parsedPepe.Look, evCtx.SvgBuilder)
		if err != nil {
			return &errors.StoreError{Problem: "Error! Failed to convert pepe data to SVG!", Keys: []string{pepeId.Text(10)}, Err: err}
		}
		svgData.Lcb = int64(currentBlock)
		res = svgData
	}

	log.Printf("SvgUpdate: trig.Block: %d, current block: %d, pepe: %s\n",
		trig.Block, currentBlock, pepeId.Text(10))

	evCtx.EntityBuf.ChangeEntity(dataKey, data.ReplaceIfNewer(dataKey, res))

	return nil
}

