package manus

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testAPIKey = "test-api-key"

func TestNewClient_DefaultConfig(t *testing.T) {
	apiKey := os.Getenv(ManusAPIKeyEnv)
	if apiKey == "" {
		t.Skip("MANUS_API_KEY not set, skipping test")
	}

	client := NewClient(apiKey)
	assert.NotNil(t, client)
	assert.NotNil(t, client.restyClient)
	assert.Equal(t, apiKey, client.apiKey)
}

func TestNewClient_WithCustomBaseURL(t *testing.T) {
	customBaseURL := "https://custom.api.example.com"

	client := NewClient(testAPIKey, WithBaseURL(customBaseURL))
	assert.NotNil(t, client)
}

func TestNewClient_WithCustomTimeout(t *testing.T) {
	customTimeout := 60 * time.Second

	client := NewClient(testAPIKey, WithTimeout(customTimeout))
	assert.NotNil(t, client)
}

func TestNewClient_WithCustomRetryConfig(t *testing.T) {
	client := NewClient(testAPIKey,
		WithRetryCount(5),
		WithRetryWaitTime(2*time.Second, 10*time.Second),
	)
	assert.NotNil(t, client)
}

func TestNewClient_WithMultipleOptions(t *testing.T) {
	client := NewClient(testAPIKey,
		WithBaseURL("https://custom.api.example.com"),
		WithTimeout(60*time.Second),
		WithRetryCount(5),
		WithRetryWaitTime(2*time.Second, 10*time.Second),
	)
	assert.NotNil(t, client)
}

func TestNewClient_WithCustomHeader(t *testing.T) {
	client := NewClient(testAPIKey,
		WithHeader("X-Custom-Header", "custom-value"),
	)
	assert.NotNil(t, client)
}

func TestNewClient_WithMultipleCustomHeaders(t *testing.T) {
	headers := map[string]string{
		"X-Request-ID": "request-123",
		"X-User-Agent": "my-app/1.0",
	}

	client := NewClient(testAPIKey,
		WithHeaders(headers),
	)
	assert.NotNil(t, client)
}

func TestNewClient_EmptyAPIKey(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty API key")
		}
	}()

	NewClient("")
}
