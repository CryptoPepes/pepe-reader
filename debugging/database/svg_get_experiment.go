package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"cryptopepe.io/cryptopepe-reader/datastoring"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
)

func main() {

	dc := datastoring.NewDataStoreClient()

	ctx := context.Background()

	svgData := &convert.PepeSVGData{}
	err := dc.Get(ctx, datastore.NameKey("svg", "1", nil), svgData)
	if err != nil {
		log.Fatalln("Failed SVG get!", err)
	}

	log.Println("SVG result: ", string(svgData.Svg))
}
