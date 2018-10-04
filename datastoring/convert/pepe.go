package convert

import (
	pepeSpec "cryptopepe.io/cryptopepe-reader/pepe"
	"cryptopepe.io/cryptopepe-svg/builder/look"
	"log"
	"cloud.google.com/go/datastore"
	"strings"
	"cryptopepe.io/cryptopepe-reader/bio-gen"
	"time"
)

func PepeToPepeData(pepe *pepeSpec.Pepe, bioGen *bio_gen.BioGenerator) *PepeData {
	data := new(PepeData)
	data.Master = strings.ToLower(pepe.Master.Hex())
	data.Genotype = pepe.Genotype[0].Text(16) + pepe.Genotype[1].Text(16)
	//convert to int64 for datastore compability, bits do not change and can be converted back.
	data.CanCozyAgain = int64(pepe.CanCozyAgain)
	data.Generation = int64(pepe.Generation)
	data.Father = PepeIdToString(pepe.Father)
	data.Mother = PepeIdToString(pepe.Mother)
	data.PepeName = nameToUTF8(pepe.PepeName[:])
	//Store in int16, no need to check sign in database query
	data.CoolDownIndex = int16(pepe.CoolDownIndex)
	// Also store the full look of the pepe.
	dna := pepeSpec.PepeDNA(pepe.Genotype)
	data.Look = *(&dna).ParsePepeDNA()
	bio := bioGen.ConvertDnaToBio(&dna)
	data.BioTitle = bio.Title
	data.BioDescription = bio.Description
	data.PepeState = "normal"
	return data
}

// Parses the name bytes, and strips off 0 bytes (both sides)
func nameToUTF8(name []byte) string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Found an invalid name, defaulted to 'pepe'.")
		}
	}()
	padL := 0
	for i := 0; i < 32; i++ {
		if name[i] != 0 {
			break
		}
		padL++
	}
	padR := 32
	for i := 31; i > padL; i-- {
		if name[i] != 0 {
			break
		}
		padR--
	}
	return string(name[padL:padR])
}

type PepeData struct {
	Master         string        `datastore:"master"`
	Genotype       string        `datastore:"genotype,noindex"`
	CanCozyAgain   int64         `datastore:"can_cozy_again"`
	Generation     int64         `datastore:"gen"`
	Father         string        `datastore:"father"`
	Mother         string        `datastore:"mother"`
	PepeName       string        `datastore:"name,noindex"`
	CoolDownIndex  int16         `datastore:"cool_down_index"`
	Look           look.PepeLook `datastore:"look"`
	BioTitle       string  		 `datastore:"bio_title,noindex"`
	BioDescription string  		 `datastore:"bio_description,noindex"`

	// State of the pepe:
	// "cozy": in active cozy auction
	// "cozy_expired": in expired cozy auction
	// "sale": in sale auction
	// "sale_expired": in expired sale auction
	// "normal": default state, in owner wallet, can breed
	// "cozy_cooldown": active cooldown, in owner wallet, bus has to wait for cozy time.
	PepeState      string        `datastore:"pepe_state"`

	// Optional, may be null in the dataset.
	// If nil, then not in active auction (Note, if not nil, auction may be elapsed)
	SaleAuction       *AuctionData  `datastore:"sale_auction"`
	CozyAuction       *AuctionData  `datastore:"cozy_auction"`

	Lcb int64 `datastore:"lcb,noindex"`
}

// Interface, si
func (data *PepeData) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(data, ps)
}
func (data *PepeData) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(data)
}

func (data *PepeData) LastChangedBlockNr() uint64 {
	return uint64(data.Lcb)
}

func (data *PepeData) DeleteMe() bool {
	return false
}

// Cleans up old auction data, and update variables in general.
// softUpdate: if true, the Lcb will be incremented if an update is needed
func (data *PepeData) Update(softUpdate bool) {
	dirty := false

	now := time.Now().Unix()
	// Check if pepe state matches current auction state, if there's any
	if data.SaleAuction != nil {
		if data.SaleAuction.IsExpired(now) {
			if data.PepeState != "sale_expired" {
				data.PepeState = "sale_expired"
				dirty = true
			}
		} else {
			if data.PepeState != "sale" {
				data.PepeState = "sale"
				dirty = true
			}
		}
	}

	// Check if pepe state matches current auction state, if there's any
	if data.CozyAuction != nil {
		if data.CozyAuction.IsExpired(now) {
			if data.PepeState != "cozy_expired" {
				data.PepeState = "cozy_expired"
				dirty = true
			}
		} else {
			if data.PepeState != "cozy" {
				data.PepeState = "cozy"
				dirty = true
			}
		}
	}

	// If not in an auction, check if the pepe state matches the cooldown state
	if data.SaleAuction == nil && data.CozyAuction == nil {
		if data.CanCozyAgain < now {
			if data.PepeState != "normal" {
				data.PepeState = "normal"
				dirty = true
			}
		} else {
			if data.PepeState != "cozy_cooldown" {
				data.PepeState = "cozy_cooldown"
				dirty = true
			}
		}
	}

	// On soft update, don't change the Lcb, all we want is the state to be updated, for internal use.
	if !softUpdate && dirty {
		// If some data was changed: update the block number by 1, to force an update
		data.Lcb++
	}
}
