package slack

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpContract interface {
	Post(url string, body io.Reader) ([]byte, error)
}

type jsonHttpClient struct {
	http http.Client
}

func (h jsonHttpClient) Post(url string, body io.Reader) ([]byte, error) {
	resp, err := h.http.Post(url, "application/json", body)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

type HttpClient struct {
	Client HttpContract
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: jsonHttpClient{},
	}
}

func (h *HttpClient) Send(url string, message Message) ([]byte, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return data, err
	}

	return h.Client.Post(url, bytes.NewBuffer(data))
}
