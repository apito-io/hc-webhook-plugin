package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-webhook-plugin] Starting Webhook plugin...")
	plugin := sdk.Init("hc-webhook-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("WebhookStatus", "Webhook plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getWebhookStatus",
		sdk.ComplexObjectField("Get webhook plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-webhook-plugin] Plugin ready")
	plugin.Serve()
}
