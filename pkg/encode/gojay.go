package encode

import (
	"github.com/francoispqt/gojay"
	"logstash-go-client"
)

type GojayEncoder struct{}

func NewGojayEncoder() *GojayEncoder {
	return &GojayEncoder{}
}

func (g GojayEncoder) Marshal(fields logstash.Fields) ([]byte, error) {
	fieldsInBytes, err := gojay.MarshalJSONObject(fields)
	if err != nil {
		return nil, err
	}
	return fieldsInBytes, nil
}
