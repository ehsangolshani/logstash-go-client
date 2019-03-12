package logstash

import (
	"github.com/francoispqt/gojay"
)

type GojayJsonEncoder struct{}

func NewGojayJsonEncoder() *GojayJsonEncoder {
	return &GojayJsonEncoder{}
}

func (g GojayJsonEncoder) Marshal(sendItem SendItem) ([]byte, error) {
	fieldsInBytes, err := gojay.MarshalJSONObject(sendItem)
	if err != nil {
		return nil, err
	}

	return fieldsInBytes, nil
}
