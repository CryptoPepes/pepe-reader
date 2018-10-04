package convert

import (
	"math/big"
	"cloud.google.com/go/datastore"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"fmt"
)

func PepeIdToString(id *big.Int) string {
	return id.Text(10)
}

func PepeIdToPaddedString(id *big.Int) string {
	// Pad with 10 0's, so that alphabetic sorting works nicely.
	return fmt.Sprintf("%010s", id.Text(10))
}

func PepeIdToKey(id *big.Int) *datastore.Key {
	return datastore.NameKey("pepe", PepeIdToPaddedString(id), nil)
}

func PepeIdToSVGKey(id *big.Int) *datastore.Key {
	return datastore.NameKey("svg", PepeIdToPaddedString(id), nil)
}

func AddressToUserKey(address *common.Address) *datastore.Key {
	// Note the lowercase: although Geth supports EIP-55, some other systems may not use it.
	// Hence, just use lowercase.
	return datastore.NameKey("user", strings.ToLower(address.Hex()), nil)
}
