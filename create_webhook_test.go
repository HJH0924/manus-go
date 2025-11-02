package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebhook(t *testing.T) {
	tests := []struct {
		name string
		req  *CreateWebhookRequest
	}{
		{
			name: "webhook with http url",
			req: &CreateWebhookRequest{
				URL: "https://example.com/webhook",
			},
		},
		{
			name: "webhook with custom domain",
			req: &CreateWebhookRequest{
				URL: "https://myapp.example.com/api/webhooks/manus",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(os.Getenv(ManusAPIKeyEnv))
			resp, err := client.CreateWebhook(context.Background(), tt.req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.WebhookID)
			t.Logf("Created webhook: %s for URL: %s", resp.WebhookID, tt.req.URL)
		})
	}
}
