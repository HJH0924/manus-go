package manus

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTask(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a task to delete
	createResp, err := client.CreateTask(&CreateTaskRequest{
		Prompt:   "Test task for DeleteTask",
		TaskMode: TaskModeChat,
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	taskID := createResp.TaskID
	t.Logf("Created task: %s", taskID)

	// Now delete the task
	deleteResp, err := client.DeleteTask(taskID)
	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
	assert.Equal(t, taskID, deleteResp.ID)
	assert.Equal(t, "task.deleted", deleteResp.Object)
	assert.True(t, deleteResp.Deleted)
	t.Logf("Deleted task: %s", taskID)
}

func TestDeleteTask_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.DeleteTask("nonexistent_task_id")
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
