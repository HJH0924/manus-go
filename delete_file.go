package manus

import (
	"context"
	"fmt"
)

// DeleteFileResponse represents the response from deleting a file.
type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// DeleteFile deletes a file by ID.
// This removes both the file record and its associated content from S3 storage.
// Reference: https://open.manus.ai/docs/api-reference/delete-file
func (c *Client) DeleteFile(ctx context.Context, fileID string) (*DeleteFileResponse, error) {
	if fileID == "" {
		return nil, fmt.Errorf("fileID is required")
	}

	var result DeleteFileResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		Delete("/v1/files/" + fileID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
