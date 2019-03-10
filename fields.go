package logstash

import (
	"github.com/francoispqt/gojay"
	"time"
)

type Fields map[string]interface{}

func (f Fields) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range f {
		enc.AddInterfaceKeyOmitEmpty(k, v)
	}
}

func (f Fields) IsNil() bool {
	return f == nil
}

func (f Fields) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var value interface{}
	err := dec.Interface(&value)
	if err != nil {
		return err
	}
	f[k] = value
	return nil
}

func (f Fields) NKeys() int {
	return 0
}

func (f Fields) WithTimestamp() {
	f["@timestamp"] = time.Now().Format(time.RFC3339)
}

func (f Fields) WithVersion(version string) {
	f["@version"] = version
}

func (f Fields) AddField(name string, value interface{}) {
	f[name] = value
}
