package textanalytics

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	marshaler   = &jsonpb.Marshaler{}
	unmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)

type Client struct {
	client   *http.Client
	endpoint string
	key      string
}

func New(client *http.Client, endpoint, key string) *Client {
	return &Client{
		client:   client,
		endpoint: endpoint,
		key:      key,
	}
}

func (c *Client) do(ctx context.Context, req, resp proto.Message) error {
	uri := c.endpoint + "/text/analytics/v2.0/languages"
	s, err := marshaler.MarshalToString(req)
	if err != nil {
		return fmt.Errorf("marshal request: %v", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(s))
	if err != nil {
		return fmt.Errorf("create http request: %v", err)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Ocp-Apim-Subscription-Key", c.key)
	httpReq.WithContext(ctx)

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("http request: %v", err)
	}
	defer httpResp.Body.Close()

	b, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %v", err)
	}

	err = unmarshaler.Unmarshal(bytes.NewReader(b), resp)
	if err != nil {
		return fmt.Errorf("unmarshal response: %v", err)
	}

	return nil
}
