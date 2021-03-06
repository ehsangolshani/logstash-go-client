package logstash

import (
	"bytes"
	"github.com/pkg/errors"
	"net/http"
)

type HttpClient struct {
	Host        string
	HttpClient  *http.Client
	JsonEncoder JsonEncoder
}

func NewHttpClient(host string, httpClient *http.Client) *HttpClient {
	return &HttpClient{Host: host, HttpClient: httpClient, JsonEncoder: NewGojayJsonEncoder()}

}

func (c *HttpClient) SetJsonEncoder(encoder JsonEncoder) {
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
