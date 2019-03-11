package http

import (
	"bytes"
	"errors"
	"logstash-go-client"
	"logstash-go-client/pkg/encode"
	"net/http"
)

type Client struct {
	Host        string
	HttpClient  *http.Client
	JsonEncoder encode.JsonEncoder
}

func NewClient(host string, httpClient *http.Client) *Client {
	return &Client{Host: host, HttpClient: httpClient, JsonEncoder: encode.NewGojayEncoder()}
}

func (c *Client) SetJsonEncoder(encoder encode.JsonEncoder) {
	c.JsonEncoder = encoder
}

func (c Client) Send(sendingItem logstash.SendItem) error {
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

func (c Client) SendBytes(messageInBytes []byte) error {
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
