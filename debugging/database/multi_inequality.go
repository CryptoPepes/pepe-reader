package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"cryptopepe.io/cryptopepe-reader/datastoring"
	"fmt"
	"google.golang.org/api/iterator"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
)

func main() {

	dc := datastoring.NewDataStoreClient()

	ctx := context.Background()

	q := datastore.NewQuery("pepe").
		Offset(3).
		Filter("gen >", 2).
		Filter("birth_time >=", 4)

	for iter := dc.Run(ctx, q); ; {
		var pepe convert.PepeData
		key, err := iter.Next(&pepe)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Failed to get next element!", key, err)
			return
		}
		fmt.Printf("Key=%v\nPepe=%#v\n\n", key, pepe)
		fmt.Println("name: ", pepe.PepeName)
	}
}
