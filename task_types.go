package manus

// Attachment represents a file/image attachment for tasks.
// It supports three types: File ID Attachment, URL Attachment, and Base64 Data Attachment.
type Attachment struct {
	Filename string `json:"filename"`
	// For File ID Attachment
	FileID string `json:"file_id,omitempty"`
	// For URL Attachment
	URL string `json:"url,omitempty"`
	// For Base64 Data Attachment
	FileData string `json:"fileData,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

// NewFileIDAttachment creates a new attachment using a file ID.
func NewFileIDAttachment(filename, fileID string) *Attachment {
	if filename == "" {
		panic("filename is required")
	}

	if fileID == "" {
		panic("fileID is required")
	}

	return &Attachment{
		Filename: filename,
		FileID:   fileID,
	}
}

// NewURLAttachment creates a new attachment using a URL.
func NewURLAttachment(filename, url, mimeType string) *Attachment {
	if filename == "" {
		panic("filename is required")
	}

	if url == "" {
		panic("url is required")
	}

	return &Attachment{
		Filename: filename,
		URL:      url,
		MimeType: mimeType,
	}
}

// NewBase64Attachment creates a new attachment using base64 encoded data.
func NewBase64Attachment(filename, fileData string) *Attachment {
	if filename == "" {
		panic("filename is required")
	}

	if fileData == "" {
		panic("fileData is required")
	}

	return &Attachment{
		Filename: filename,
		FileData: fileData,
	}
}

// TaskMode represents the task mode.
type TaskMode string

const (
	// TaskModeChat represents chat mode.
	TaskModeChat TaskMode = "chat"
	// TaskModeAdaptive represents adaptive mode.
	TaskModeAdaptive TaskMode = "adaptive"
	// TaskModeAgent represents agent mode.
	TaskModeAgent TaskMode = "agent"
)

// AgentProfile represents the agent profile.
type AgentProfile string

const (
	// AgentProfileSpeed represents speed profile.
	AgentProfileSpeed AgentProfile = "speed"
	// AgentProfileQuality represents quality profile.
	AgentProfileQuality AgentProfile = "quality"
)

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	// TaskStatusPending indicates the task is pending.
	TaskStatusPending TaskStatus = "pending"
	// TaskStatusRunning indicates the task is running.
	TaskStatusRunning TaskStatus = "running"
	// TaskStatusCompleted indicates the task is completed.
	TaskStatusCompleted TaskStatus = "completed"
	// TaskStatusFailed indicates the task has failed.
	TaskStatusFailed TaskStatus = "failed"
)

// TaskMetaData represents custom key-value pairs for task metadata.
type TaskMetaData map[string]any

// MessageRole represents the role of a message sender.
type MessageRole string

const (
	// MessageRoleUser represents a user message.
	MessageRoleUser MessageRole = "user"
	// MessageRoleAssistant represents an assistant message.
	MessageRoleAssistant MessageRole = "assistant"
)

// MessageContentType represents the type of message content.
type MessageContentType string

const (
	// MessageContentTypeText represents a text message.
	MessageContentTypeText MessageContentType = "output_text"
	// MessageContentTypeFile represents a file message.
	MessageContentTypeFile MessageContentType = "output_file"
)

// MessageContent represents the content of a message.
type MessageContent struct {
	Type     MessageContentType `json:"type,omitempty"`
	Text     string             `json:"text,omitempty"`
	FileURL  string             `json:"fileUrl,omitempty"`
	FileName string             `json:"fileName,omitempty"`
	MimeType string             `json:"mimeType,omitempty"`
}

// Message represents a message in the task output.
type Message struct {
	ID      string           `json:"id"`
	Status  string           `json:"status"`
	Role    MessageRole      `json:"role"`
	Type    string           `json:"type"`
	Content []MessageContent `json:"content"`
}

// Task represents a task in the Manus API.
// Reference: https://open.manus.ai/docs/api-reference/get-task
type Task struct {
	ID                string       `json:"id"`
	Object            string       `json:"object"`
	CreatedAt         string       `json:"created_at"` // Unix timestamp as string
	UpdatedAt         string       `json:"updated_at"` // Unix timestamp as string
	Status            TaskStatus   `json:"status"`
	Error             string       `json:"error,omitempty"`
	IncompleteDetails string       `json:"incomplete_details,omitempty"`
	Instructions      string       `json:"instructions,omitempty"`
	MaxOutputTokens   int          `json:"max_output_tokens,omitempty"`
	Model             string       `json:"model"`
	Metadata          TaskMetaData `json:"metadata"`
	Output            []Message    `json:"output"`
	Locale            string       `json:"locale,omitempty"`
	CreditUsage       int          `json:"credit_usage,omitempty"`
}
