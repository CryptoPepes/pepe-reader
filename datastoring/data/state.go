package data

import (
	"cloud.google.com/go/datastore"
	"sync"
	"context"
	storeErrors "cryptopepe.io/cryptopepe-reader/datastoring/errors"
	"time"
	"log"
	"errors"
	"cryptopepe.io/cryptopepe-reader/datastoring/convert"
	"google.golang.org/api/iterator"
)

type ContentStateVars struct {
	Lcb int64 `datastore:"lcb"`
}

// Ignores all other struct fields, only loads lcb.
func (data *ContentStateVars) Load(ps []datastore.Property) error {
	lcbIndex := -1
	for i := 0; i < len(ps); i++ {
		prop := ps[i]
		if prop.Name == "lcb" {
			lcbIndex = i
			break
		}
	}
	if lcbIndex < 0 {
		return errors.New("datastore has invalid format, missing lcb property")
	}
	// Load as usual, but only the lcb variable.
	return datastore.LoadStruct(data, ps[lcbIndex:lcbIndex+1])
}

func (data *ContentStateVars) Save() ([]datastore.Property, error) {
	return nil, errors.New("this type of content should not be saved")
}

func (data *ContentStateVars) LastChangedBlockNr() uint64 {
	return uint64(data.Lcb)
}

func (data *ContentStateVars) DeleteMe() bool {
	return false
}


type ContentState interface {

	datastore.PropertyLoadSaver

	//The block number of the last update.
	// (Update as in update of the program state,
	//  not just incremental smartcontract changes,
	//  re-orgs can cause this too!)
	LastChangedBlockNr() uint64

	//If this entity should be deleted from the database.
	DeleteMe() bool
}

type DeletableContent struct {
	Lcb uint64
}

func NewDeletableContent(lastChange uint64) *DeletableContent {
	res := DeletableContent{Lcb: lastChange}
	return &res
}


// Ignores all other struct fields, only loads lcb.
func (d *DeletableContent) Load(ps []datastore.Property) error {
	return errors.New("this type of content should not be loaded")
}

func (d *DeletableContent) Save() ([]datastore.Property, error) {
	return nil, errors.New("this type of content should not be saved")
}

func (d *DeletableContent) LastChangedBlockNr() uint64 {
	return d.Lcb
}

func (d *DeletableContent) DeleteMe() bool {
	return true
}

type EntityState struct {
	sync.Mutex

	Key *datastore.Key

	//Contents of the entity,
	// which should be able to provide the block number corresponding to the last change.
	Content ContentState

}

type EntityBuffer struct {
	sync.Mutex

	entities map[string]*EntityState

	//the datastore backing this buffer
	dc *datastore.Client
}

func NewEntityBuffer(dc *datastore.Client) *EntityBuffer {
	eb := new(EntityBuffer)
	eb.entities = make(map[string]*EntityState)
	eb.dc = dc
	return eb
}

// Wait till the entity buffers shrinks to under 10000 entries. Out-of-memory pre-caution.
func (eb *EntityBuffer) WaitReady() {
	for i := 0; i < 20; i++ {
		if len(eb.entities) < 10000 {
			return
		}

		<- time.After(time.Second * 30)
	}
	log.Println("Warning: timed-out wait-ready on event-buffer. Ignoring it.")
}

// prevState may be nil if not present
// returns (write, next entity state).
// If write is false, no changes are made.
// If next state is nil, entity is removed from buffer, but not deleted from datastore.
// Mark an entity as deleted to have it be deleted in the next update round.
type StateChangeFn func(prevState *EntityState) (bool, *EntityState)


type QueryMaker func() *datastore.Query

func (eb *EntityBuffer) UpdatePepesData() {
	go func() {
		if err := eb.updateFromIter(func() *datastore.Query {
			// find elapsed auctions
			return datastore.NewQuery("pepe").
				Filter("pepe_state =", "sale").
				Filter("sale_auction.end_time <", time.Now().Unix())
		}); err != nil {
			log.Println("Failed to update SaleAuction data.", err)
		}
	}()
	go func() {
		if err := eb.updateFromIter(func() *datastore.Query {
			// find elapsed auctions
			return datastore.NewQuery("pepe").
				Filter("pepe_state =", "cozy").
				Filter("cozy_auction.end_time <", time.Now().Unix())
		}); err != nil {
			log.Println("Failed to update CozyAuction data.", err)
		}
	}()
}

func (eb *EntityBuffer) updateFromIter(queryMaker QueryMaker) error {

	// Execute a small keys-only query first, which is much more economical.
	keyOnlyCheck := queryMaker().Limit(1).KeysOnly()
	checkIter := eb.dc.Run(getUpdateCtx(), keyOnlyCheck)
	_, err := checkIter.Next(nil)
	if err == iterator.Done {
		log.Println("No pepes with old data to update in datastore!")
		return nil
	}
	if err != nil {
		log.Println("Failed to check datastore for old pepe data!")
		return err
	}

	// Run the full query only when we know there are results to update.
	query := queryMaker()

	iter := eb.dc.Run(getUpdateCtx(), query)

	for {
		dsPepe := convert.PepeData{}
		dataKey, err := iter.Next(&dsPepe)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Could not get pepe from query iterator. ", err)
			continue
		}
		dsPepe.Update(false)

		eb.ChangeEntity(dataKey, ReplaceIfNewer(dataKey, &dsPepe))
	}

	return nil
}

// Get entity state
func (eb *EntityBuffer) GetEntity(key *datastore.Key) *EntityState {
	return eb.entities[key.String()]
}

// Atomic change of entity state, can also delete entities (nextState = nil).
func (eb *EntityBuffer) ChangeEntity(key *datastore.Key, fn StateChangeFn) *EntityState {
	eb.Lock()
	defer eb.Unlock()

	keyStr := key.String()
	ent := eb.entities[keyStr]
	write, next := fn(ent)
	if !write {
		//return previous entity when there was no write.
		return ent
	}

	if next == nil {
		delete(eb.entities, keyStr)
	} else {
		eb.entities[keyStr] = next
	}

	return next
}

// Google Datastore accepts batches of 500 max.
// But go for a bit less, partial failure is a better worst case.
const MaxUpdateSize = 200

func (eb *EntityBuffer) UpdateAll() {
	eb.Lock()
	defer eb.Unlock()


	keys := make([]*datastore.Key, 0, len(eb.entities))
	keyStrs := make([]string, 0, len(eb.entities))
	values := make([]*EntityState, 0, len(eb.entities))
	for keyStr, value := range eb.entities {
		keys = append(keys, value.Key)
		keyStrs = append(keyStrs, keyStr)
		values = append(values, value)
	}

	// Update in batches
	total := len(keys)
	if total == 0 {
		log.Println("Skipping unnecessary update, no entities to update.")
		return
	}
	for i := 0; i < total; i += MaxUpdateSize {
		end := i + MaxUpdateSize
		if end > total {
			end = total
		}
		if err := eb.updateWithKeys(keys[i:end], keyStrs[i:end], values[i:end]); err != nil {
			log.Printf("Warning! Failed to update datastore with keys! Key window: [%d, %d). Err: %s\n", i, end, err)
		}
	}
}

func getUpdateCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	return ctx
}

func (eb *EntityBuffer) updateWithKeys(keys []*datastore.Key, keyStrs []string, values []*EntityState) error {

	if len(keys) == 0 {
		log.Println("Skipping update, no keys to update!")
	}

	//unsafe, no lock. Use UpdateAll() !!!
	log.Printf("Updating datastore from buffer! Entity count: %d\n", len(keys))

	//Check if the datastore has newer values
	dsContents := make([]ContentStateVars, len(keys))
	if err := handleMultiErr(eb.dc.GetMulti(getUpdateCtx(), keys, dsContents)); err != nil {
		return err
	}

	//filter out the contents that are not newer than those in the datastore
	keyStrs2, values2 := retainNewerContents(keyStrs, values, dsContents)

	// Remove any entities from the buffer that are too old.
	// (i.e. datastore has more recent data on it)
	tooOld := mergeDiff(keyStrs, keyStrs2)
	for _, oldKey := range tooOld {
		delete(eb.entities, oldKey)
	}

	if len(keyStrs2) == 0 {
		log.Println("Skipping update, all keys were too old in comparison to DB, removed them!")
	}

	putKeys, putValues, deleteKeys := splitUpdateModes(values2)

	// Overwrite old entities
	if _, err := eb.dc.PutMulti(getUpdateCtx(), putKeys, putValues); err != nil {
		return &storeErrors.StoreError{Problem: "Error! Failed to put entities!", Keys: getKeyStrs(putKeys), Err: err}
	}

	// Delete old entities
	if err := eb.dc.DeleteMulti(getUpdateCtx(), deleteKeys); err != nil {
		return &storeErrors.StoreError{Problem: "Error! Failed to delete entities!", Keys: getKeyStrs(deleteKeys), Err: err}
	}

	// Do not clean up buffer for entities that were put into the database
	// Just call UpdateDatastore() again,
	//  and it will check if the datastore was updated completely,
	//  removing any stale buffer contents.

	return nil
}

func getKeyStrs(keys []*datastore.Key) (keyStrs []string) {
	keyStrs = make([]string, 0, len(keys))
	for _, key := range keys {
		keyStrs = append(keyStrs, key.String())
	}
	return keyStrs
}

func splitUpdateModes(ents []*EntityState) (putKeys []*datastore.Key, putValues []ContentState, deleteKeys []*datastore.Key) {
	putKeys = make([]*datastore.Key, 0, len(ents))
	putValues = make([]ContentState, 0, len(ents))
	deleteKeys = make([]*datastore.Key, 0, len(ents))
	for _, ent := range ents {
		if ent.Content.DeleteMe() {
			deleteKeys = append(deleteKeys, ent.Key)
		} else {
			putKeys = append(putKeys, ent.Key)
			putValues = append(putValues, ent.Content)
		}
	}
	return putKeys, putValues, deleteKeys
}

// Returns the values in all that are not present in sub
func mergeDiff(all []string, sub []string) (diff []string) {
	diff = make([]string, 0, len(all) - len(sub))
	for i, j := 0, 0; j < len(all); {
		if i < len(sub) && all[j] == sub[i] {
			i++
			j++
		} else {
			diff = append(diff, all[j])
			j++
		}
	}
	return diff
}

// Checks the individual errors in a multi-error
func handleMultiErr(err error) error {
	if err == nil {
		return nil
	}

	if multiErr, ok := err.(datastore.MultiError); ok {
		for _, elErr := range multiErr {
			if elErr == nil || elErr == datastore.ErrNoSuchEntity {
				//save to ignore, no entity == nil
				continue
			} else {
				log.Println(elErr)
				return multiErr
			}
		}
	} else {
		return err
	}
	return nil
}

func retainNewerContents(keyStrs []string,
	valuesA []*EntityState, valuesB []ContentStateVars) (
		newKeyStrs []string,
		newValuesA []*EntityState){

	// only update the datastore entries that are actually outdated.
	newKeyStrs = make([]string, 0, len(keyStrs))
	newValuesA = make([]*EntityState, 0, len(valuesA))

	// Loop through all values and see which are newer.
	// Retain all of valuesA that are newer than their corresponding value in B
	for i, keyStr := range keyStrs {
		contentA := valuesA[i]
		contentB := valuesB[i]
		//if contentB is default struct, then block number will be 0.
		// So absent DB entries will always be updated.
		if contentB.LastChangedBlockNr() < contentA.Content.LastChangedBlockNr() {
			newKeyStrs = append(newKeyStrs, keyStr)
			newValuesA = append(newValuesA, contentA)
		}
	}

	return newKeyStrs, newValuesA
}
