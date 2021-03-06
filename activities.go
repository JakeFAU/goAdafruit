package goadafruit

import (
	"fmt"
	"time"
)

type Activities struct {
	Username  string    `json:"username"`
	Owner     Owner     `json:"owner"`
	ID        int       `json:"id"`
	Action    string    `json:"action"`
	Model     string    `json:"model"`
	Data      string    `json:"data"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type ActivitiesService struct {
	client *Client
}

func (s *ActivitiesService) All() ([]*Activities, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/activities", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	activities := make([]*Activities, 0)
	resp, err := s.client.Do(req, &activities)
	if err != nil {
		return nil, resp, err
	}

	return activities, resp, nil
}

func (s *ActivitiesService) ActivitiesByType(atype *string) ([]*Activities, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/activities/%v", s.client.Username, *atype)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	activities := make([]*Activities, 0)
	resp, err := s.client.Do(req, &activities)
	if err != nil {
		return nil, resp, err
	}

	return activities, resp, nil
}
