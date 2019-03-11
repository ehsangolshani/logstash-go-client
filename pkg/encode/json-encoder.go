package encode

import "logstash-go-client"

type JsonEncoder interface {
	Marshal(fields logstash.Fields) ([]byte, error)
}
