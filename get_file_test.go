package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a file
	createResp, err := client.CreateFile(ctx, &CreateFileRequest{
		Filename: "test_get_file.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	fileID := createResp.ID
	t.Logf("Created file: %s", fileID)

	// Upload file content
	err = client.UploadFileContent(ctx, createResp.UploadURL, []byte("Test content for GetFile"))
	assert.NoError(t, err)

	// Now get the file details
	file, err := client.GetFile(ctx, fileID)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, fileID, file.ID)
	assert.Equal(t, "file", file.Object)
	assert.Equal(t, "test_get_file.txt", file.Filename)
	assert.NotEmpty(t, file.Status)
	assert.NotEmpty(t, file.CreatedAt)
	t.Logf("File: ID=%s, Filename=%s, Status=%s, CreatedAt=%s", file.ID, file.Filename, file.Status, file.CreatedAt)
}

func TestGetFile_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.GetFile(context.Background(), "nonexistent_file_id")
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
