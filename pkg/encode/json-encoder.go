package encode

import "github.com/ehsangolshani/logstash-go-client"

type JsonEncoder interface {
	Marshal(sendItem logstash.SendItem) ([]byte, error)
}
