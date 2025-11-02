package manus

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteWebhook(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	// First create a webhook to delete
	createResp, err := client.CreateWebhook(ctx, &CreateWebhookRequest{
		URL: "https://example.com/webhook-to-delete",
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	webhookID := createResp.WebhookID
	t.Logf("Created webhook: %s", webhookID)

	// Now delete the webhook
	deleteResp, err := client.DeleteWebhook(ctx, webhookID)
	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
	assert.True(t, deleteResp.Success)
	t.Logf("Deleted webhook: %s", webhookID)
}

func TestDeleteWebhook_NotFound(t *testing.T) {
	client := NewClient(os.Getenv(ManusAPIKeyEnv))

	_, err := client.DeleteWebhook(context.Background(), "nonexistent_webhook_id")
	assert.Error(t, err)
	t.Logf("Expected error: %v", err)
}
