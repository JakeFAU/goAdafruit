package goadafruit

import "fmt"

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Feeds       []Feed `json:"feeds"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GroupService struct {
	client *Client
}

// All returns all Groups for the current account.
func (s *GroupService) All() ([]*Group, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates Feed slice
	groups := make([]*Group, 0)
	resp, err := s.client.Do(req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// Create makes a new Group and either returns a new Group instance or an error.
func (s *GroupService) Create(g *Group) (*Group, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, g)
	if rerr != nil {
		return nil, nil, rerr
	}

	var group Group
	resp, err := s.client.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return &group, resp, nil
}

// Get returns the Group record identified by the given ID
func (s *GroupService) Get(id interface{}) (*Group, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var group Group
	resp, err := s.client.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return &group, resp, nil
}

// Update takes an ID and a Group record, updates it, and returns a new Group
// instance or an error.
func (s *GroupService) Update(id interface{}, group *Group) (*Group, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("PATCH", path, group)
	if rerr != nil {
		return nil, nil, rerr
	}

	var updatedGroup Group
	resp, err := s.client.Do(req, &updatedGroup)
	if err != nil {
		return nil, resp, err
	}

	return &updatedGroup, resp, nil
}

// Delete the Group identified by the given ID.
func (s *GroupService) Delete(id interface{}) (*Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v", s.client.Username, id)

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

// Create a feed in a group
func (s *GroupService) CreateFeedInGroup(id interface{}, f *Feed) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v/feeds", s.client.Username, id)

	req, rerr := s.client.NewRequest("POST", path, f)
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

// Add a feed to a group
func (s GroupService) AddFeedToGroup(id interface{}, f *Feed) (*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v/add", s.client.Username, id)

	req, rerr := s.client.NewRequest("POST", path, f)
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

// Remove a feed from a group
func (s GroupService) RemoveFeedFromGroup(id interface{}, f *Feed) (*Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v/remove", s.client.Username, id)

	req, rerr := s.client.NewRequest("POST", path, f)
	if rerr != nil {
		return nil, rerr
	}

	var feed Feed
	resp, err := s.client.Do(req, &feed)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List all group feeds
func (s GroupService) ListGroupFeeds(id interface{}) ([]*Feed, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/groups/%v/feeds", s.client.Username, id)

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
