package logstash

import (
	"fmt"
	"github.com/francoispqt/gojay"
	"time"
)

type Fields map[string]interface{}

func (f Fields) MarshalJSONObject(enc *gojay.Encoder) {
	a := Entry{}
	fmt.Println(a)

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

func (f Fields) WithTimestamp() *Entry {
	timestamp := time.Now()
	return &Entry{Fields: f, Timestamp: &timestamp, timeFormat: time.RFC3339}
}

func (f Fields) WithTimeFormat(timeFormat string) *Entry {
	return &Entry{Fields: f, timeFormat: timeFormat}
}

func (f Fields) WithVersion(version string) *Entry {
	return &Entry{Fields: f, Version: version}
}

func (f Fields) WithField(name string, value interface{}) *Entry {
	f[name] = value
	return &Entry{Fields: f}
}

func (f Fields) WithFields(data map[string]interface{}) *Entry {
	for name, value := range data {
		f[name] = value
	}
	return &Entry{Fields: f}
}
