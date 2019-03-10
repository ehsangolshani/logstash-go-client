package logstash

type Sender interface {
	SendFields(fields Fields) error
	SendBytes(messageInBytes []byte) error
}
