package goadafruit

import (
	"fmt"
	"path"
	"time"
)

type FeedCreation struct {
	Name          string      `json:"name,omitempty"`
	Key           string      `json:"key,omitempty"`
	Description   string      `json:"description,omitempty"`
	UnitType      interface{} `json:"unit_type,omitempty"`
	UnitSymbol    interface{} `json:"unit_symbol,omitempty"`
	History       bool        `json:"history,omitempty"`
	Visibility    string      `json:"visibility,omitempty"`
	License       interface{} `json:"license,omitempty"`
	StatusNotify  bool        `json:"status_notify,omitempty"`
	StatusTimeout int         `json:"status_timeout,omitempty"`
	Enabled       bool        `json:"enabled,omitempty"`
}

type Feed struct {
	Username             string              `json:"username,omitempty"`
	Owner                Owner               `json:"owner,omitempty"`
	ID                   int                 `json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Description          string              `json:"description,omitempty"`
	License              interface{}         `json:"license,omitempty"`
	History              bool                `json:"history,omitempty"`
	Enabled              bool                `json:"enabled,omitempty"`
	Visibility           string              `json:"visibility,omitempty"`
	UnitType             interface{}         `json:"unit_type,omitempty"`
	UnitSymbol           interface{}         `json:"unit_symbol,omitempty"`
	LastValue            string              `json:"last_value,omitempty"`
	CreatedAt            interface{}         `json:"created_at,omitempty"`
	UpdatedAt            interface{}         `json:"updated_at,omitempty"`
	StatusNotify         bool                `json:"status_notify,omitempty"`
	StatusTimeout        int                 `json:"status_timeout,omitempty"`
	Status               string              `json:"status,omitempty"`
	Key                  string              `json:"key,omitempty"`
	Writable             bool                `json:"writable,omitempty"`
	Group                Group               `json:"group,omitempty"`
	Groups               []Groups            `json:"groups,omitempty"`
	FeedWebhookReceivers []interface{}       `json:"feed_webhook_receivers,omitempty"`
	FeedStatusChanges    []FeedStatusChanges `json:"feed_status_changes,omitempty"`
}

type Groups struct {
	ID     int    `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}
type FeedStatusChanges struct {
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	FromStatus  string      `json:"from_status,omitempty"`
	ToStatus    string      `json:"to_status,omitempty"`
	EmailSent   interface{} `json:"email_sent,omitempty"`
	EmailSentTo interface{} `json:"email_sent_to,omitempty"`
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
func (s *FeedService) Create(feed *FeedCreation) (*FeedCreation, *Response, error) {
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

// Create takes a Feed record, creates it, and returns the updated record or an error.
func (s *FeedService) CreateInGroup(feed *FeedCreation, groupName string) (*FeedCreation, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/feeds?group_key=%v", s.client.Username, groupName)

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
