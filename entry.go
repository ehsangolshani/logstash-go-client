package logstash

import (
	"fmt"
	"github.com/francoispqt/gojay"
	"time"
)

type Entry struct {
	Fields     Fields
	Version    string
	Timestamp  *time.Time
	timeFormat string
}

func (e Entry) MarshalJSONObject(enc *gojay.Encoder) {
	a := Entry{}
	fmt.Println(a)

	for k, v := range e.Fields {
		enc.AddInterfaceKeyOmitEmpty(k, v)
	}
	enc.StringKeyNullEmpty("@version", e.Version)
	enc.TimeKey("@timestamp", e.Timestamp, e.timeFormat)
}

func (e Entry) IsNil() bool {
	return &e == nil
}

func (e Entry) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&e.Version)
	case "name":
		return dec.Time(e.Timestamp, e.timeFormat)
	default:
		var value interface{}
		err := dec.Interface(&value)
		if err != nil {
			return err
		}
		e.Fields[k] = value
		return nil
	}
}

func (e Entry) NKeys() int {
	return 0
}

func (e Entry) WithTimestamp() *Entry {
	timestamp := time.Now()
	e.Timestamp = &timestamp
	e.timeFormat = time.RFC3339
	return &e
}

func (e Entry) WithTimeFormat(timeFormat string) *Entry {
	e.timeFormat = time.RFC3339
	return &e
}

func (e Entry) WithVersion(version string) *Entry {
	e.Version = version
	return &e
}

func (e Entry) WithField(name string, value interface{}) *Entry {
	e.Fields[name] = value
	return &e
}

func (e Entry) WithFields(data map[string]interface{}) *Entry {
	for name, value := range data {
		e.Fields[name] = value
	}
	return &e
}
