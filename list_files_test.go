package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFiles(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	resp, err := client.ListFiles(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "list", resp.Object)
	t.Logf("Found %d files (max 10)", len(resp.Data))

	for i, file := range resp.Data {
		assert.NotEmpty(t, file.ID)
		assert.Equal(t, "file", file.Object)
		assert.NotEmpty(t, file.Filename)
		assert.NotEmpty(t, file.Status)
		assert.NotEmpty(t, file.CreatedAt)
		t.Logf("File %d: ID=%s, Filename=%s, Status=%s", i+1, file.ID, file.Filename, file.Status)
	}
}
