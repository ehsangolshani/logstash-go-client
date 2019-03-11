package encode

import (
	"github.com/francoispqt/gojay"
	"logstash-go-client"
)

type GojayEncoder struct{}

func NewGojayEncoder() *GojayEncoder {
	return &GojayEncoder{}
}

func (g GojayEncoder) Marshal(sendItem logstash.SendItem) ([]byte, error) {
	fieldsInBytes, err := gojay.MarshalJSONObject(sendItem)
	if err != nil {
		return nil, err
	}

	return fieldsInBytes, nil
}
