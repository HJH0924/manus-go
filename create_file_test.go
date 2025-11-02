package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		content  []byte
	}{
		{
			name:     "simple text file",
			filename: "test.txt",
			content:  []byte("Hello, World!"),
		},
		{
			name:     "json file",
			filename: "data.json",
			content:  []byte(`{"key": "value"}`),
		},
		{
			name:     "markdown file",
			filename: "readme.md",
			content:  []byte("# Test Markdown\n\nThis is a test."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := NewClient(os.Getenv(ManusAPIKeyEnv))

			// Step 1: Create file record and get upload URL
			createResp, err := client.CreateFile(ctx, &CreateFileRequest{
				Filename: tt.filename,
			})
			assert.NoError(t, err)
			assert.NotNil(t, createResp)
			assert.NotEmpty(t, createResp.ID)
			assert.Equal(t, "file", createResp.Object)
			assert.Equal(t, tt.filename, createResp.Filename)
			assert.Equal(t, "pending", createResp.Status)
			assert.NotEmpty(t, createResp.UploadURL)
			assert.NotEmpty(t, createResp.CreatedAt)
			t.Logf("FileID: %s, Filename: %s", createResp.ID, createResp.Filename)
			t.Logf("UploadURL: %s", createResp.UploadURL)
			t.Logf("UploadExpiresAt: %s", createResp.UploadExpiresAt)

			// Step 2: Upload file content to the presigned URL
			err = client.UploadFileContent(ctx, createResp.UploadURL, tt.content)
			assert.NoError(t, err)
			t.Logf("File content uploaded successfully")
		})
	}
}

func TestCreateFile_OnlyFilename(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// Test creating a file record without uploading content
	resp, err := client.CreateFile(context.Background(), &CreateFileRequest{
		Filename: "placeholder.txt",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ID)
	assert.NotEmpty(t, resp.UploadURL)
	t.Logf("Created file record: %s, UploadURL available for later upload", resp.ID)
}
