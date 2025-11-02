package manus

import (
	"context"
	"fmt"
)

// UpdateTaskRequest represents the request to update a task.
type UpdateTaskRequest struct {
	Title                   string `json:"title,omitempty"`
	EnableShared            bool   `json:"enableShared,omitempty"`
	EnableVisibleInTaskList bool   `json:"enableVisibleInTaskList,omitempty"`
}

// UpdateTaskResponse represents the response from updating a task.
type UpdateTaskResponse struct {
	TaskID    string `json:"task_id"`
	TaskTitle string `json:"task_title"`
	TaskURL   string `json:"task_url"`
	ShareURL  string `json:"share_url,omitempty"`
}

// UpdateTask updates an existing task.
// Reference: https://open.manus.ai/docs/api-reference/update-task
func (c *Client) UpdateTask(ctx context.Context, taskID string, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	if taskID == "" {
		return nil, fmt.Errorf("taskID is required")
	}

	var result UpdateTaskResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Put("/v1/tasks/" + taskID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
