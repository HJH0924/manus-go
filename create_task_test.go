package manus

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name string
		req  *CreateTaskRequest
	}{
		{
			name: "simple chat task",
			req: &CreateTaskRequest{
				Prompt:   "What's the weather like today in Singapore?",
				TaskMode: TaskModeChat,
			},
		},
		{
			name: "task with shareable link",
			req: &CreateTaskRequest{
				Prompt:              "Execute uptime",
				TaskMode:            TaskModeAgent,
				CreateShareableLink: true,
			},
		},
		{
			name: "task with locale",
			req: &CreateTaskRequest{
				Prompt:   "今天新加坡的天气如何？",
				TaskMode: TaskModeChat,
				Locale:   "zh-CN",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(os.Getenv(ManusAPIKeyEnv))
			resp, err := client.CreateTask(tt.req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.TaskID)
			assert.NotEmpty(t, resp.TaskTitle)
			assert.NotEmpty(t, resp.TaskURL)
			t.Logf("TaskID: %s, Title: %s, URL: %s", resp.TaskID, resp.TaskTitle, resp.TaskURL)

			if tt.req.CreateShareableLink {
				assert.NotEmpty(t, resp.ShareURL)
				t.Logf("ShareURL: %s", resp.ShareURL)
			}
		})
	}
}
