package manus

import "context"

// ListFilesResponse represents the response from listing files.
type ListFilesResponse struct {
	Object string `json:"object"`
	Data   []File `json:"data"`
}

// ListFiles retrieves the 10 most recently uploaded files.
// Reference: https://open.manus.ai/docs/api-reference/list-files
func (c *Client) ListFiles(ctx context.Context) (*ListFilesResponse, error) {
	var result ListFilesResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		Get("/v1/files")
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
