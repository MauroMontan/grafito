package grafito

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Query struct {
	Name      string
	Arguments map[string]interface{}
	Fields    []string
}

type graphqlRequest struct {
	Query string `json:"query"`
}

// Build genera el string GraphQL del Query
func (q Query) Build() string {
	var buf bytes.Buffer
	buf.WriteString(q.Name)

	if len(q.Arguments) > 0 {
		buf.WriteString("(")
		first := true
		for k, v := range q.Arguments {
			if !first {
				buf.WriteString(", ")
			}
			first = false
			switch val := v.(type) {
			case string:
				buf.WriteString(fmt.Sprintf("%s: \"%s\"", k, val))
			default:
				buf.WriteString(fmt.Sprintf("%s: %v", k, val))
			}
		}
		buf.WriteString(")")
	}

	if len(q.Fields) > 0 {
		buf.WriteString(" { ")
		for _, f := range q.Fields {
			buf.WriteString(f + " ")
		}
		buf.WriteString("}")
	}

	return buf.String()
}

func (client *Client) AddHeader(key string, value string) *Client {
	client.Header.Add(key, value)
	return client
}

func (client *Client) doPost(ctx context.Context, _payload io.Reader) *http.Request {

	req, httpErr := http.NewRequestWithContext(ctx, "POST", client.url, _payload)

	req.Header = client.Header

	if httpErr != nil {
		panic(httpErr)
	}

	return req
}

func (client *Client) RunQuery(ctx context.Context, q Query, dest any) error {

	fullQuery := fmt.Sprintf("{ %s }", q.Build())

	return client.run(ctx, fullQuery, dest)

}

func (client *Client) run(ctx context.Context, query string, dest any) error {

	graphqlRequest := graphqlRequest{Query: query}
	body, _ := json.Marshal(graphqlRequest)

	_payload := bytes.NewReader(body)

	req := client.doPost(ctx, _payload)

	fmt.Printf("req.Header: %v\n", req.Header)

	resp, requestErr := http.DefaultClient.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	err := json.Unmarshal(data, &dest)

	if err != nil {
		return err
	}

	return nil

}
