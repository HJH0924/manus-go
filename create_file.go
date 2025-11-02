package manus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// CreateFileRequest represents the request to create a file.
// Reference: https://open.manus.ai/docs/api-reference/create-file
type CreateFileRequest struct {
	Filename string `json:"filename"` // Required: Name of the file to upload
}

// CreateFileResponse represents the response from creating a file.
type CreateFileResponse struct {
	ID              string `json:"id"`
	Object          string `json:"object"` // Always "file"
	Filename        string `json:"filename"`
	Status          string `json:"status"`            // Initial status is "pending"
	UploadURL       string `json:"upload_url"`        // Presigned S3 URL for uploading file content
	UploadExpiresAt string `json:"upload_expires_at"` // ISO 8601 timestamp when upload URL expires
	CreatedAt       string `json:"created_at"`        // ISO 8601 timestamp
}

// CreateFile creates a file record in the Manus API and returns a presigned URL for uploading.
// This is step 1 of the file upload process. Use UploadFileContent() for step 2.
// Reference: https://open.manus.ai/docs/api-reference/create-file
func (c *Client) CreateFile(ctx context.Context, req *CreateFileRequest) (*CreateFileResponse, error) {
	if req.Filename == "" {
		return nil, fmt.Errorf("filename is required")
	}

	var file CreateFileResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&file).
		Post("/v1/files")
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &file, nil
}

// UploadFileContent uploads file content to the presigned S3 URL.
// This is step 2 of the file upload process, after calling CreateFile().
// The uploadURL is obtained from CreateFileResponse.UploadURL.
func (c *Client) UploadFileContent(ctx context.Context, uploadURL string, content []byte) error {
	if uploadURL == "" {
		return fmt.Errorf("uploadURL is required")
	}

	if len(content) == 0 {
		return fmt.Errorf("content is required")
	}

	// Use a new resty client for uploading to S3 (without API headers)
	client := resty.New()

	resp, err := client.R().
		SetContext(ctx).
		SetBody(content).
		Put(uploadURL)
	if err != nil {
		return fmt.Errorf("failed to upload file content: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}
