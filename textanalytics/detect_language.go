package textanalytics

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (c *Client) BatchDetectLanguage(ctx context.Context, req *BatchDetectLanguageRequest) (*BatchDetectLanguageResponse, error) {
	resp := &BatchDetectLanguageResponse{}
	err := c.do(ctx, detectLanguagePath, req, resp)
	return resp, err
}

func (c *Client) DetectLanguage(ctx context.Context, text string) ([]*DetectedLanguage, error) {
	id := uuid.New().String()
	resp, err := c.BatchDetectLanguage(ctx, &BatchDetectLanguageRequest{
		Documents: []*BatchDetectLanguageRequest_Document{
			{
				Id:   id,
				Text: text,
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
	return resp.Documents[0].DetectLanguages, nil
}
