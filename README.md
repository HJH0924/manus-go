# Manus Go SDK

Go SDK for the [Manus API](https://open.manus.ai/docs), built with [resty](https://github.com/go-resty/resty).

## Installation

```bash
go get github.com/HJH0924/manus-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/HJH0924/manus-go"
)

func main() {
    // Initialize client with your API key
    client := manus.NewClient("your-api-key-here")
    ctx := context.Background()

    // Create a new task
    taskReq := &manus.CreateTaskRequest{
        Prompt:   "Hello, world!",
        TaskMode: manus.TaskModeChat,
    }

    task, err := client.CreateTask(ctx, taskReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created task: %s\n", task.TaskID)
}
```

## Client Configuration

You can configure the client with various options:

```go
import "time"

client := manus.NewClient(
    "your-api-key-here",
    manus.WithBaseURL("https://custom-url.com"),                        // Custom base URL
    manus.WithTimeout(60 * time.Second),                                // Custom timeout
    manus.WithRetryCount(5),                                            // Custom retry count
    manus.WithRetryWaitTime(2*time.Second, 10*time.Second),            // Custom retry wait time
    manus.WithHeader("Custom-Header", "value"),                         // Add custom header
    manus.WithHeaders(map[string]string{"Header1": "value1"}),         // Add multiple headers
)
```

## API Reference

### Tasks

#### Create Task

```go
ctx := context.Background()

taskReq := &manus.CreateTaskRequest{
    Prompt:              "Your prompt here",
    TaskMode:            manus.TaskModeAgent,
    AgentProfile:        manus.AgentProfileQuality,
    Attachments:         []manus.Attachment{
        *manus.NewFileIDAttachment("document.pdf", "file-id-1"),
        *manus.NewURLAttachment("image.jpg", "https://example.com/image.jpg", "image/jpeg"),
    },
    CreateShareableLink: true,
    Locale:              "en-US",
}

task, err := client.CreateTask(ctx, taskReq)
```

#### Get Tasks List

```go
ctx := context.Background()

// List tasks with filters
params := &manus.GetTasksRequest{
    Limit:   10,
    Order:   "desc",
    OrderBy: "created_at",
    Status:  []manus.TaskStatus{manus.TaskStatusCompleted},
    Query:   "search term",
}

tasks, err := client.GetTasks(ctx, params)

// Pagination
if tasks.HasMore {
    nextPage := &manus.GetTasksRequest{
        After: tasks.LastID,
        Limit: 10,
    }
    moreTasks, err := client.GetTasks(ctx, nextPage)
}
```

#### Get Task

```go
ctx := context.Background()

task, err := client.GetTask(ctx, "task-id")
```

#### Update Task

```go
ctx := context.Background()

updateReq := &manus.UpdateTaskRequest{
    Title:                   "New Title",
    EnableShared:            true,
    EnableVisibleInTaskList: true,
}

result, err := client.UpdateTask(ctx, "task-id", updateReq)
```

#### Delete Task

```go
ctx := context.Background()

result, err := client.DeleteTask(ctx, "task-id")
if result.Deleted {
    fmt.Println("Task deleted successfully")
}
```

### Files

#### Create and Upload File (Two-Step Process)

File upload is a two-step process:
1. Create a file record and get a presigned upload URL
2. Upload the file content to the presigned URL

```go
ctx := context.Background()

// Step 1: Create file record
fileReq := &manus.CreateFileRequest{
    Filename: "document.pdf",
}

fileResp, err := client.CreateFile(ctx, fileReq)
if err != nil {
    log.Fatal(err)
}

// Step 2: Upload file content to S3
content, err := os.ReadFile("path/to/document.pdf")
if err != nil {
    log.Fatal(err)
}

err = client.UploadFileContent(ctx, fileResp.UploadURL, content)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("File uploaded: %s\n", fileResp.ID)
```

#### List Files

```go
ctx := context.Background()

files, err := client.ListFiles(ctx)
for _, file := range files.Data {
    fmt.Printf("File: %s (%s)\n", file.Filename, file.Status)
}
```

#### Get File

```go
ctx := context.Background()

file, err := client.GetFile(ctx, "file-id")
```

#### Delete File

```go
ctx := context.Background()

result, err := client.DeleteFile(ctx, "file-id")
if result.Deleted {
    fmt.Println("File deleted successfully")
}
```

### Webhooks

#### Create Webhook

```go
ctx := context.Background()

webhookReq := &manus.CreateWebhookRequest{
    URL: "https://your-webhook-url.com/webhook",
}

webhook, err := client.CreateWebhook(ctx, webhookReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Webhook created: %s\n", webhook.WebhookID)
```

#### Delete Webhook

```go
ctx := context.Background()

result, err := client.DeleteWebhook(ctx, "webhook-id")
if result.Success {
    fmt.Println("Webhook deleted successfully")
}
```

## Error Handling

The SDK returns `*manus.APIError` for API errors:

```go
ctx := context.Background()

task, err := client.GetTask(ctx, "task-id")
if err != nil {
    if apiErr, ok := err.(*manus.APIError); ok {
        fmt.Printf("API Error: %s (HTTP %d, Code %d)\n",
            apiErr.Message, apiErr.StatusCode, apiErr.Code)
    } else {
        fmt.Printf("Network Error: %v\n", err)
    }
}
```

## License

See [LICENSE](LICENSE) file for details.
