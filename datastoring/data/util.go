package data

import "cloud.google.com/go/datastore"

// The generic data replacer: only update data when the update is more recently retrieved from the blockchain.
func ReplaceIfNewer(key *datastore.Key, replacement ContentState) func(prevState *EntityState) (bool, *EntityState) {

	if replacement == nil {
		panic("Replacement for "+key.String()+" is nil!")
	}

	return func(prevState *EntityState) (bool, *EntityState) {
		if prevState == nil ||
			prevState.Content.LastChangedBlockNr() < replacement.LastChangedBlockNr() {

			next := new(EntityState)
			next.Key = key
			next.Content = replacement

			return true, next
		}
		return false, nil
	}
}