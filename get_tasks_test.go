package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string
		params *GetTasksRequest
	}{
		{
			name:   "list all tasks",
			params: nil,
		},
		{
			name: "with limit",
			params: &GetTasksRequest{
				Limit: 5,
			},
		},
		{
			name: "with status filter",
			params: &GetTasksRequest{
				Status: []TaskStatus{TaskStatusCompleted},
			},
		},
		{
			name: "with order",
			params: &GetTasksRequest{
				Limit:   10,
				Order:   "asc",
				OrderBy: "created_at",
			},
		},
		{
			name: "with search query",
			params: &GetTasksRequest{
				Query: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(os.Getenv(ManusAPIKeyEnv))
			resp, err := client.GetTasks(ctx, tt.params)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "list", resp.Object)
			t.Logf("Found %d tasks, HasMore: %v", len(resp.Data), resp.HasMore)

			if len(resp.Data) > 0 {
				t.Logf("FirstID: %s, LastID: %s", resp.FirstID, resp.LastID)
			}
		})
	}
}

func TestGetTasks_Pagination(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First page
	resp1, err := client.GetTasks(ctx, &GetTasksRequest{
		Limit: 2,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp1)
	t.Logf("Page 1: %d tasks", len(resp1.Data))

	// Second page if available
	if resp1.HasMore {
		resp2, err := client.GetTasks(ctx, &GetTasksRequest{
			After: resp1.LastID,
			Limit: 2,
		})
		assert.NoError(t, err)
		assert.NotNil(t, resp2)
		t.Logf("Page 2: %d tasks", len(resp2.Data))
	}
}
