package manus

import (
	"context"
	"fmt"
)

// DeleteTaskResponse represents the response from deleting a task.
type DeleteTaskResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// DeleteTask deletes a task by ID.
// Reference: https://open.manus.ai/docs/api-reference/delete-task
func (c *Client) DeleteTask(ctx context.Context, taskID string) (*DeleteTaskResponse, error) {
	if taskID == "" {
		return nil, fmt.Errorf("taskID is required")
	}

	var result DeleteTaskResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		Delete("/v1/tasks/" + taskID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
