package manus

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a task
	createResp, err := client.CreateTask(&CreateTaskRequest{
		Prompt:   "Test task for GetTask",
		TaskMode: TaskModeChat,
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	taskID := createResp.TaskID
	t.Logf("Created task: %s", taskID)

	// Now get the task details
	task, err := client.GetTask(taskID)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, taskID, task.ID)
	assert.Equal(t, "task", task.Object)
	assert.NotEmpty(t, task.Model)
	assert.NotEmpty(t, task.CreatedAt)
	t.Logf("Task Status: %s, Model: %s, Created: %s", task.Status, task.Model, task.CreatedAt)
	t.Logf("Output messages: %d", len(task.Output))
}

func TestGetTask_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.GetTask("nonexistent_task_id")
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
