package logstash

import (
	"bytes"
	"github.com/pkg/errors"
	"logstash-go-client/pkg/encode"
	"net/http"
)

type HttpClient struct {
	Host        string
	HttpClient  *http.Client
	JsonEncoder encode.JsonEncoder
}

func NewClient(host string, httpClient *http.Client) *HttpClient {
	return &HttpClient{Host: host, HttpClient: httpClient, JsonEncoder: encode.NewGojayEncoder()}
}

func (c *HttpClient) SetJsonEncoder(encoder encode.JsonEncoder) {
	c.JsonEncoder = encoder
}

func (c HttpClient) Send(sendingItem SendItem) error {
	fieldsInBytes, err := c.JsonEncoder.Marshal(sendingItem)
	if err != nil {
		return err
	}

	err = c.SendBytes(fieldsInBytes)
	if err != nil {
		return err
	}

	return nil
}

func (c HttpClient) SendBytes(messageInBytes []byte) error {
	request, err := http.NewRequest("POST", c.Host, bytes.NewBuffer(messageInBytes))
	if err != nil {

		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.HttpClient.Do(request)
	if err != nil {

		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("received non-200 response status code")
	}

	return nil
}
