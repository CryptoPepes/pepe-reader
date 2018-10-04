package convert

import (
	"cloud.google.com/go/datastore"
)

func UserToUserData(usernameBytes [32]byte) *UserData {
	data := new(UserData)

	data.UserName = string(usernameBytes[:])

	return data
}

// User data is really simple: just store the username,
//  the address is already used and retrievable as the data entry key.
type UserData struct {
	UserName string `datastore:"username"`

	Lcb int64 `datastore:"lcb,noindex"`
}

// Interface, si
func (data *UserData) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(data, ps)
}
func (data *UserData) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(data)
}

func (data *UserData) LastChangedBlockNr() uint64 {
	return uint64(data.Lcb)
}

func (data *UserData) DeleteMe() bool {
	return false
}

