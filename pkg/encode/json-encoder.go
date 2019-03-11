package encode

import logstash "logstash-go-client"

type JsonEncoder interface {
	Marshal(sendItem logstash.SendItem) ([]byte, error)
}
