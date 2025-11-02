package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFile(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a file to delete
	createResp, err := client.CreateFile(ctx, &CreateFileRequest{
		Filename: "test_delete_file.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	fileID := createResp.ID
	t.Logf("Created file: %s", fileID)

	// Upload file content
	err = client.UploadFileContent(ctx, createResp.UploadURL, []byte("Test content for DeleteFile"))
	assert.NoError(t, err)

	// Now delete the file
	deleteResp, err := client.DeleteFile(ctx, fileID)
	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
	assert.Equal(t, fileID, deleteResp.ID)
	assert.Equal(t, "file.deleted", deleteResp.Object) // According to the API documentation, the object should be "file.deleted", but actually it is "file"
	assert.True(t, deleteResp.Deleted)
	t.Logf("Deleted file: %s", fileID)
}

func TestDeleteFile_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.DeleteFile(context.Background(), "nonexistent_file_id")
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
