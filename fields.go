package logstash

import (
	"github.com/francoispqt/gojay"
	"time"
)

type Fields struct {
	Data       map[string]interface{}
	Version    string
	Timestamp  *time.Time
	timeFormat string
}

func (f Fields) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range f.Data {
		enc.AddInterfaceKeyOmitEmpty(k, v)
	}
	enc.StringKeyNullEmpty("@version", f.Version)
	enc.TimeKey("@timestamp", f.Timestamp, f.timeFormat)
}

func (f Fields) IsNil() bool {
	return &f == nil
}

func (f Fields) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&f.Version)
	case "name":
		return dec.Time(f.Timestamp, f.timeFormat)
	default:
		var value interface{}
		err := dec.Interface(&value)
		if err != nil {
			return err
		}
		f.Data[k] = value
		return nil
	}
}

func (f Fields) NKeys() int {
	return 0
}

func (f *Fields) WithTimestamp() *Fields {
	timestamp := time.Now()
	f.Timestamp = &timestamp
	f.timeFormat = time.RFC3339
	return f
}

func (f *Fields) WithTimeFormat(timeFormat string) *Fields {
	f.timeFormat = timeFormat
	return f
}

func (f *Fields) WithVersion(version string) *Fields {
	f.Version = version
	return f
}

func (f *Fields) WithField(name string, value interface{}) *Fields {
	f.Data[name] = value
	return f
}

func (f *Fields) WithFields(data map[string]interface{}) *Fields {
	for name, value := range data {
		f.Data[name] = value
	}
	return f
}
