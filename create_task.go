package manus

// CreateTaskRequest represents the request to create a task.
type CreateTaskRequest struct {
	Prompt              string       `json:"prompt"`
	Attachments         []Attachment `json:"attachments,omitempty"`
	TaskMode            TaskMode     `json:"taskMode,omitempty"`
	Connectors          []string     `json:"connectors,omitempty"`
	HideInTaskList      bool         `json:"hideInTaskList,omitempty"`
	CreateShareableLink bool         `json:"createShareableLink,omitempty"`
	TaskID              string       `json:"taskId,omitempty"`
	AgentProfile        AgentProfile `json:"agentProfile,omitempty"`
	Locale              string       `json:"locale,omitempty"`
}

// CreateTaskResponse represents the response from creating a task.
type CreateTaskResponse struct {
	TaskID    string `json:"task_id"`
	TaskTitle string `json:"task_title"`
	TaskURL   string `json:"task_url"`
	ShareURL  string `json:"share_url,omitempty"`
}

// CreateTask creates a new AI task with custom parameters and attachments.
// Reference: https://open.manus.ai/docs/api-reference/create-task
func (c *Client) CreateTask(req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var result CreateTaskResponse

	resp, err := c.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/v1/tasks")
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
