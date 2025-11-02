// Package manus provides a Go SDK for the Manus API.
package manus

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client is the main client for interacting with the Manus API.
type Client struct {
	restyClient *resty.Client
	apiKey      string
}

// ClientOption is a function type for setting client options.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the client.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.restyClient.SetBaseURL(baseURL)
	}
}

// WithTimeout sets a custom timeout for the client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.restyClient.SetTimeout(timeout)
	}
}

// WithRetryCount sets the number of retries for failed requests.
func WithRetryCount(count int) ClientOption {
	return func(c *Client) {
		c.restyClient.SetRetryCount(count)
	}
}

// WithRetryWaitTime sets the wait time and max wait time between retries.
func WithRetryWaitTime(waitTime, maxWaitTime time.Duration) ClientOption {
	return func(c *Client) {
		if waitTime > 0 {
			c.restyClient.SetRetryWaitTime(waitTime)
		}

		if maxWaitTime > 0 {
			c.restyClient.SetRetryMaxWaitTime(maxWaitTime)
		}
	}
}

// WithHeader adds a custom header to all requests.
func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.restyClient.SetHeader(key, value)
	}
}

// WithHeaders adds multiple custom headers to all requests.
func WithHeaders(headers map[string]string) ClientOption {
	return func(c *Client) {
		c.restyClient.SetHeaders(headers)
	}
}

// NewClient creates a new Manus API client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	if apiKey == "" {
		panic("MANUS_API_KEY is required")
	}

	restyClient := resty.New()
	restyClient.SetTimeout(DefaultTimeout)
	restyClient.SetRetryCount(DefaultRetryCount)
	restyClient.SetRetryWaitTime(DefaultRetryWaitTime)
	restyClient.SetRetryMaxWaitTime(DefaultRetryMaxWaitTime)
	restyClient.SetBaseURL(DefaultBaseURL)

	client := &Client{
		restyClient: restyClient,
		apiKey:      apiKey,
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	// Set default headers
	restyClient.SetHeader(HeaderAPIKey, apiKey)
	restyClient.SetHeader(HeaderContentType, ContentTypeJSON)

	return client
}

// handleResponse handles the API response and converts errors if needed.
func (c *Client) handleResponse(resp *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if resp.IsError() {
		apiErr := &APIError{
			StatusCode: resp.StatusCode(),
		}

		// Try to parse error response
		if err := json.Unmarshal(resp.Body(), apiErr); err != nil {
			// If parsing fails, use status text
			apiErr.Message = resp.Status()
		}

		return apiErr
	}

	return nil
}
