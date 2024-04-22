package grafito

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	url    string
	Header http.Header
}

func NewClient(url string) *Client {
	return &Client{
		url:    url,
		Header: http.Header{},
	}

}

func (client *Client) AddHeader(key string, value string) *Client {
	client.Header.Add(key, value)
	return client
}

func (client *Client) doPost(_payload io.Reader) *http.Request {

	req, httpErr := http.NewRequest("POST", client.url, _payload)

	req.Header = client.Header

	if httpErr != nil {
		panic(httpErr)
	}

	return req
}
func (client *Client) Query(payload string, object any) error {

	_payload := strings.NewReader(payload)

	req := client.doPost(_payload)

	fmt.Printf("req.Header: %v\n", req.Header)

	resp, requestErr := http.DefaultClient.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	err := json.Unmarshal(body, &object)

	if err != nil {
		return err
	}

	return nil

}
