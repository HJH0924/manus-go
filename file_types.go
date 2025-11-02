package manus

// FileStatus represents the status of a file.
type FileStatus string

const (
	// FileStatusPending indicates the file is pending upload.
	FileStatusPending FileStatus = "pending"
	// FileStatusUploaded indicates the file has been uploaded.
	FileStatusUploaded FileStatus = "uploaded"
	// FileStatusDeleted indicates the file has been deleted.
	FileStatusDeleted FileStatus = "deleted"
)

// File represents a file in the Manus API.
type File struct {
	ID        string     `json:"id"`
	Object    string     `json:"object"`
	Filename  string     `json:"filename"`
	Status    FileStatus `json:"status"`
	CreatedAt string     `json:"created_at"`
}
