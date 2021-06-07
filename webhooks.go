package goadafruit

import "fmt"

type WebhookData struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	FeedID    int    `json:"feed_id"`
	FeedKey   string `json:"feed_key"`
	CreatedAt string `json:"created_at"`
	Location  struct {
	} `json:"location"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	Ele          float64 `json:"ele"`
	CreatedEpoch int     `json:"created_epoch"`
	Expiration   string  `json:"expiration"`
}

type WebhookService struct {
	client *Client
}

func (s *WebhookService) SendDataViaWebhook(token string, d *WebhookData) (*WebhookData, *Response, error) {

	path := fmt.Sprintf("/api/v2/webhooks/feed/%v", token)

	req, rerr := s.client.NewRequest("POST", path, d)
	if rerr != nil {
		return nil, nil, rerr
	}

	var data WebhookData
	resp, err := s.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data, resp, nil
}

func (s *WebhookService) SendArbitraryDataViaWebhook(token string, d *WebhookData) (*WebhookData, *Response, error) {

	path := fmt.Sprintf("/api/v2/webhooks/feed/%v/raw", token)

	req, rerr := s.client.NewRequest("POST", path, d)
	if rerr != nil {
		return nil, nil, rerr
	}

	var data WebhookData
	resp, err := s.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data, resp, nil
}

func (s *WebhookService) SendNotificationViaWebhook(token string, d *WebhookData) (*WebhookData, *Response, error) {

	path := fmt.Sprintf("/api/v2/webhooks/feed/%v/notify", token)

	req, rerr := s.client.NewRequest("POST", path, d)
	if rerr != nil {
		return nil, nil, rerr
	}

	var data WebhookData
	resp, err := s.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data, resp, nil
}
