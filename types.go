package manus

import "fmt"

// APIError represents an error response from the API.
type APIError struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Details    []string `json:"details"`
	StatusCode int      `json:"-"` // HTTP status code (not from JSON)
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("API error (HTTP %d, code %d): %s", e.StatusCode, e.Code, e.Message)
	}

	return fmt.Sprintf("API error (code %d): %s", e.Code, e.Message)
}
