package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendDataViaWebhook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/webhooks/feed/test-token",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{
			"id": "wh1",
			"value": "42",
			"feed_id": 0,
			"feed_key": "test-feed"
		  }`)
		})

	assert := assert.New(t)

	hook, resp, err := client.Webhook.SendDataViaWebhook("test-token", &WebhookData{ID: "wh1"})

	assert.Nil(err)
	assert.NotNil(hook)
	assert.NotNil(resp)

	assert.Equal("wh1", hook.ID)
}

func TestSendArbitraryWebhookData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/webhooks/feed/test-token/raw",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{
		"id": "wh1",
		"value": "This is non JSON raw data",
		"feed_id": 0,
		"feed_key": "test-feed"
	  }`)
		})

	assert := assert.New(t)

	hook, resp, err := client.Webhook.SendArbitraryDataViaWebhook("test-token", &WebhookData{ID: "wh1"})

	assert.Nil(err)
	assert.NotNil(hook)
	assert.NotNil(resp)

	assert.Equal("wh1", hook.ID)
}

func TestSendNotificationViaWebhook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/webhooks/feed/test-token/notify",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{
			"id": "wh1",
			"value": "ping",
			"feed_id": 0,
			"feed_key": "test-feed"
		}`)
		})

	assert := assert.New(t)

	hook, resp, err := client.Webhook.SendNotificationViaWebhook("test-token", &WebhookData{ID: "wh1"})

	assert.Nil(err)
	assert.NotNil(hook)
	assert.NotNil(resp)

	assert.Equal("wh1", hook.ID)

}
