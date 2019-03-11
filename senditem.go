package logstash

import "github.com/francoispqt/gojay"

type SendItem interface {
	WithTimestamp() *Entry
	WithTimeFormat(timeFormat string) *Entry
	WithVersion(version string) *Entry
	WithField(name string, value interface{}) *Entry
	WithFields(data map[string]interface{}) *Entry
	MarshalJSONObject(enc *gojay.Encoder)
	IsNil() bool
}
