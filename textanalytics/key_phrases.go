package textanalytics

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (c *Client) BatchKeyPhrases(ctx context.Context, req *BatchKeyPhrasesRequest) (*BatchKeyPhrasesResponse, error) {
	resp := &BatchKeyPhrasesResponse{}
	err := c.do(ctx, keyPhrasesPath, req, resp)
	return resp, err
}

func (c *Client) KeyPhrases(ctx context.Context, language, text string) ([]string, error) {
	id := uuid.New().String()
	resp, err := c.BatchKeyPhrases(ctx, &BatchKeyPhrasesRequest{
		Documents: []*BatchKeyPhrasesRequest_Document{
			{
				Language: language,
				Id:       id,
				Text:     text,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Documents) != 1 {
		return nil, fmt.Errorf("expect 1 document; got %v", len(resp.Documents))
	}
	if resp.Documents[0].GetId() != id {
		return nil, fmt.Errorf("response id %v not equal request id %v", resp.Documents[0].GetId(), id)
	}
	return resp.Documents[0].KeyPhrases, nil
}
