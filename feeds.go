package goadafruit

import (
	"fmt"
	"path"
	"time"
)

type Feed struct {
	Username             string              `json:"username"`
	Owner                Owner               `json:"owner"`
	ID                   int                 `json:"id"`
	Name                 string              `json:"name"`
	Description          string              `json:"description"`
	License              interface{}         `json:"license"`
	History              bool                `json:"history"`
	Enabled              bool                `json:"enabled"`
	Visibility           string              `json:"visibility"`
	UnitType             interface{}         `json:"unit_type"`
	UnitSymbol           interface{}         `json:"unit_symbol"`
	LastValue            string              `json:"last_value"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
	StatusNotify         bool                `json:"status_notify"`
	StatusTimeout        int                 `json:"status_timeout"`
	Status               string              `json:"status"`
	Key                  string              `json:"key"`
	Writable             bool                `json:"writable"`
	Group                Group               `json:"group"`
	Groups               []Groups            `json:"groups"`
	FeedWebhookReceivers []interface{}       `json:"feed_webhook_receivers"`
	FeedStatusChanges    []FeedStatusChanges `json:"feed_status_changes"`
}

type Groups struct {
	ID     int    `json:"id"`
	Key    string `json:"key"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
type FeedStatusChanges struct {
	CreatedAt   time.Time   `json:"created_at"`
	FromStatus  string      `json:"from_status"`
	ToStatus    string      `json:"to_status"`
	EmailSent   interface{} `json:"email_sent"`
	EmailSentTo interface{} `json:"email_sent_to"`
}

type FeedService struct {
	// CurrentFeed is the Feed used for all Data access.
	CurrentFeed *Feed

	client *Client
}

// Path generates a Feed-specific path with the given suffix.
func (s *FeedService) Path(suffix string) (string, error) {
	ferr := s.client.checkFeed()
	if ferr != nil {
		return "", ferr
	}
	return path.Join(fmt.Sprintf("api/v2/%v/feeds/%v", s.client.Username, s.CurrentFeed.Key), suffix), nil
}

// All lists all available feeds.
func (s *FeedService) All() ([]*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates Feed slice
	feeds := make([]*Feed, 0)
	resp, err := s.client.Do(req, &feeds)
	if err != nil {
		return nil, resp, err
	}

	return feeds, resp, nil
}

// Get returns the Feed record identified by the given parameter. Parameter can
// be the Feed's Name, Key, or ID.
func (s *FeedService) Get(id interface{}) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var feed Feed
	resp, err := s.client.Do(req, &feed)
	if err != nil {
		return nil, resp, err
	}

	return &feed, resp, nil
}

// Get returns the Feed record identified by the given parameter. Parameter can
// be the Feed's Name, Key, or ID.
func (s *FeedService) GetDetails(id interface{}) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds/%v/details", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var feed Feed
	resp, err := s.client.Do(req, &feed)
	if err != nil {
		return nil, resp, err
	}

	return &feed, resp, nil
}

// Create takes a Feed record, creates it, and returns the updated record or an error.
func (s *FeedService) Create(feed *Feed) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, feed)
	if rerr != nil {
		return nil, nil, rerr
	}

	resp, err := s.client.Do(req, feed)
	if err != nil {
		return nil, resp, err
	}

	return feed, resp, nil
}

// Update takes an ID and a Feed record, updates it, and returns an updated
// record instance or an error.
//
// Only the Feed Name and Description can be modified.
func (s *FeedService) Update(id interface{}, feed *Feed) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("PATCH", path, feed)
	if rerr != nil {
		return nil, nil, rerr
	}

	var updatedFeed Feed
	resp, err := s.client.Do(req, &updatedFeed)
	if err != nil {
		return nil, resp, err
	}

	return &updatedFeed, resp, nil
}

// Delete the Feed identified by the given ID.
func (s *FeedService) Delete(id interface{}) (*Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("DELETE", path, nil)
	if rerr != nil {
		return nil, rerr
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
