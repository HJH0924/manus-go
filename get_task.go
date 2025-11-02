package manus

import (
	"context"
	"fmt"
)

// GetTask retrieves a specific task by ID.
// Reference: https://open.manus.ai/docs/api-reference/get-task
func (c *Client) GetTask(ctx context.Context, taskID string) (*Task, error) {
	if taskID == "" {
		return nil, fmt.Errorf("taskID is required")
	}

	var task Task

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetResult(&task).
		Get("/v1/tasks/" + taskID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &task, nil
}
