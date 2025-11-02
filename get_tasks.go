package manus

import "fmt"

// GetTasksRequest represents the request parameters for getting tasks list.
type GetTasksRequest struct {
	After         string       `json:"after,omitempty"`
	Limit         int          `json:"limit,omitempty"`
	Order         string       `json:"order,omitempty"`
	OrderBy       string       `json:"orderBy,omitempty"`
	Query         string       `json:"query,omitempty"`
	Status        []TaskStatus `json:"status,omitempty"`
	CreatedAfter  int          `json:"createdAfter,omitempty"`
	CreatedBefore int          `json:"createdBefore,omitempty"`
}

// GetTasksResponse represents the response from getting tasks list.
type GetTasksResponse struct {
	Object  string `json:"object"`
	Data    []Task `json:"data"`
	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
	HasMore bool   `json:"has_more"`
}

// GetTasks retrieves a list of tasks with optional filtering and pagination.
// Reference: https://open.manus.ai/docs/api-reference/get-tasks
func (c *Client) GetTasks(params *GetTasksRequest) (*GetTasksResponse, error) {
	var result GetTasksResponse

	req := c.restyClient.R().
		SetResult(&result)

	if params != nil {
		if params.After != "" {
			req.SetQueryParam("after", params.After)
		}

		if params.Limit > 0 {
			req.SetQueryParam("limit", fmt.Sprintf("%d", params.Limit))
		}

		if params.Order != "" {
			req.SetQueryParam("order", params.Order)
		}

		if params.OrderBy != "" {
			req.SetQueryParam("orderBy", params.OrderBy)
		}

		if params.Query != "" {
			req.SetQueryParam("query", params.Query)
		}

		if len(params.Status) > 0 {
			for _, status := range params.Status {
				req.SetQueryParam("status", string(status))
			}
		}

		if params.CreatedAfter > 0 {
			req.SetQueryParam("createdAfter", fmt.Sprintf("%d", params.CreatedAfter))
		}

		if params.CreatedBefore > 0 {
			req.SetQueryParam("createdBefore", fmt.Sprintf("%d", params.CreatedBefore))
		}
	}

	resp, err := req.Get("/v1/tasks")
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
