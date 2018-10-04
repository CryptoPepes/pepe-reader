package datastoring

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"log"
)

func NewDataStoreClient() *datastore.Client {
	ctx := context.Background()

	// Creates a client. (empty projectID -> library will use environment var)
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}