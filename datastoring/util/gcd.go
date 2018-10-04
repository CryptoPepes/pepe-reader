package util

import (
	"cloud.google.com/go/datastore"
	"context"
)

/**
Google Cloud Datastore util functions
*/

// Returns true if the entity exists. False otherwise.
func EntityExists(dc *datastore.Client, key *datastore.Key) bool {

	res := struct{}{}

	ctx := context.Background()

	if err := dc.Get(ctx, key, &res); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return false
		} else {
			panic(err)
		}
	}
	return true
}
