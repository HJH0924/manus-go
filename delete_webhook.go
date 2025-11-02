package manus

import (
	"context"
	"fmt"
)

// DeleteWebhookResponse represents the response from deleting a webhook.
type DeleteWebhookResponse struct {
	Success bool `json:"-"`
}

// DeleteWebhook removes a previously registered webhook.
// Reference: https://open.manus.ai/docs/api-reference/delete-webhook
func (c *Client) DeleteWebhook(ctx context.Context, webhookID string) (*DeleteWebhookResponse, error) {
	if webhookID == "" {
		return nil, fmt.Errorf("webhookID is required")
	}

	resp, err := c.restyClient.R().
		SetContext(ctx).
		Delete("/v1/webhooks/" + webhookID)
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	// API returns 204 No Content on success
	return &DeleteWebhookResponse{Success: true}, nil
}
