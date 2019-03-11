package logstash

type Sender interface {
	Send(fields Fields) error
	SendBytes(messageInBytes []byte) error
}
