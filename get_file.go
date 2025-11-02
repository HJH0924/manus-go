package manus

import (
	"context"
	"fmt"
)

// GetFile retrieves a specific file by ID.
// Reference: https://open.manus.ai/docs/api-reference/get-file
func (c *Client) GetFile(ctx context.Context, fileID string) (*File, error) {
	if fileID == "" {
		return nil, fmt.Errorf("fileID is required")
	}

	var file File

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetResult(&file).
		Get("/v1/files/" + fileID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &file, nil
}
