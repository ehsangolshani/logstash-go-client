package encode

import (
	"github.com/ehsangolshani/logstash-go-client"
	"github.com/francoispqt/gojay"
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
