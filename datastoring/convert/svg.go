package convert

import (
	"cryptopepe.io/cryptopepe-svg/builder/look"
	"cryptopepe.io/cryptopepe-svg/builder"
	"bytes"
	"cloud.google.com/go/datastore"
)

func PepeToSVG(look *look.PepeLook, svgBuilder *builder.SVGBuilder) (*PepeSVGData, error) {
	data := new(PepeSVGData)
	buf := new(bytes.Buffer)
	err := svgBuilder.ConvertToSVG(buf, look)
	if err != nil {
		return nil, err
	}

	data.Svg = buf.Bytes()
	return data, nil
}

type PepeSVGData struct {
	Svg []byte `datastore:"blob,noindex"`

	Lcb int64 `datastore:"lcb,noindex"`
}

// Interface, si
func (data *PepeSVGData) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(data, ps)
}
func (data *PepeSVGData) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(data)
}

func (data *PepeSVGData) LastChangedBlockNr() uint64 {
	return uint64(data.Lcb)
}

func (data *PepeSVGData) DeleteMe() bool {
	return false
}

