package manus

import (
	"context"
	"fmt"
)

// CreateWebhookRequest represents the request to create a webhook.
type CreateWebhookRequest struct {
	URL string `json:"url"`
}

// CreateWebhookResponse represents the response from creating a webhook.
type CreateWebhookResponse struct {
	WebhookID string `json:"webhook_id"`
}

// CreateWebhook registers a webhook to receive real-time notifications.
// Reference: https://open.manus.ai/docs/api-reference/create-webhook
func (c *Client) CreateWebhook(ctx context.Context, req *CreateWebhookRequest) (*CreateWebhookResponse, error) {
	if req.URL == "" {
		return nil, fmt.Errorf("URL is required")
	}

	// API requires the URL to be wrapped in a "webhook" object
	requestBody := map[string]any{
		"webhook": map[string]string{
			"url": req.URL,
		},
	}

	var webhook CreateWebhookResponse

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetBody(requestBody).
		SetResult(&webhook).
		Post("/v1/webhooks")
	if err := c.handleResponse(resp, err); err != nil {
		return nil, err
	}

	return &webhook, nil
}
