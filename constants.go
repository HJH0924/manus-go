package manus

import "time"

// Environment variable names.
const (
	// ManusAPIKeyEnv is the environment variable name for Manus API key.
	ManusAPIKeyEnv = "MANUS_API_KEY" // #nosec G101 -- This is not a hardcoded credential, just a constant name
)

// API configuration constants.
const (
	// DefaultBaseURL is the default base URL for Manus API.
	DefaultBaseURL = "https://api.manus.ai"
	// DefaultTimeout is the default timeout for API requests.
	DefaultTimeout = 30 * time.Second
	// DefaultRetryCount is the default number of retries for failed requests.
	DefaultRetryCount = 3
	// DefaultRetryWaitTime is the default wait time between retries.
	DefaultRetryWaitTime = 1 * time.Second
	// DefaultRetryMaxWaitTime is the default maximum wait time between retries.
	DefaultRetryMaxWaitTime = 5 * time.Second
)

// API header names.
const (
	// HeaderAPIKey is the header name for API key authentication.
	HeaderAPIKey = "API_KEY"
	// HeaderContentType is the header name for content type.
	HeaderContentType = "Content-Type"
	// ContentTypeJSON is the JSON content type.
	ContentTypeJSON = "application/json"
)
