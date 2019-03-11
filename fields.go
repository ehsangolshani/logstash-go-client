package logstash

import (
	"time"
)

type Fields map[string]interface{}

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
