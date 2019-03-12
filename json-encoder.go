package logstash

type JsonEncoder interface {
	Marshal(sendItem SendItem) ([]byte, error)
}
