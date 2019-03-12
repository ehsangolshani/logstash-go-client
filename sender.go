package logstash

type Sender interface {
	Send(sendingItem SendItem) error
	SendBytes(messageInBytes []byte) error
}
