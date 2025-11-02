package manus

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTask(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a task
	createResp, err := client.CreateTask(&CreateTaskRequest{
		Prompt:   "Test task for UpdateTask",
		TaskMode: TaskModeChat,
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	taskID := createResp.TaskID
	t.Logf("Created task: %s", taskID)

	tests := []struct {
		name string
		req  *UpdateTaskRequest
	}{
		{
			name: "update title",
			req: &UpdateTaskRequest{
				Title: "Updated Task Title",
			},
		},
		{
			// This test case not work
			// return: {"code":2, "message":"proto: syntax error (line 2:3): invalid value enableShared", "details":[]}
			name: "enable shared",
			req: &UpdateTaskRequest{
				EnableShared: true,
			},
		},
		{
			// This test case not work
			// return: {"code":5, "message":"task not found or does not belong to user", "details":[]}
			name: "update visibility",
			req: &UpdateTaskRequest{
				EnableVisibleInTaskList: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.UpdateTask(taskID, tt.req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, taskID, resp.TaskID)
			assert.NotEmpty(t, resp.TaskURL)
			t.Logf("Updated task: %s, Title: %s, URL: %s", resp.TaskID, resp.TaskTitle, resp.TaskURL)

			if tt.req.EnableShared {
				assert.NotEmpty(t, resp.ShareURL)
				t.Logf("ShareURL: %s", resp.ShareURL)
			}
		})
	}
}

func TestUpdateTask_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.UpdateTask("nonexistent_task_id", &UpdateTaskRequest{
		Title: "New Title",
	})
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
