package goadafruit

import (
	"fmt"
	"time"
)

type Dashboard struct {
	Username       string      `json:"username"`
	Owner          Owner       `json:"owner"`
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Key            string      `json:"key"`
	Description    string      `json:"description"`
	Visibility     string      `json:"visibility"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	ShowHeader     bool        `json:"show_header"`
	ColorMode      string      `json:"color_mode"`
	BlockBorders   bool        `json:"block_borders"`
	HeaderImageURL interface{} `json:"header_image_url"`
}

type DashboardService struct {
	client *Client
}

func (s *DashboardService) AllDashboards() ([]*Dashboard, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	boards := make([]*Dashboard, 0)
	resp, err := s.client.Do(req, &boards)
	if err != nil {
		return nil, resp, err
	}

	return boards, resp, nil
}

func (s *DashboardService) CreateDashboard(d *Dashboard) (*Dashboard, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, d)
	if rerr != nil {
		return nil, nil, rerr
	}
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return &d, resp, nil

}

func (s *DashboardService) GetDashboard(id string) (*Dashboard, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	d := Dashboard{Name: "DB1"}

	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return &d, resp, nil

}

func (s *DashboardService) ChangeDashboard(id string, d Dashboard) (*Dashboard, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("PUT", path, d)
	if rerr != nil {
		return nil, nil, rerr
	}

	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return &d, resp, nil

}

func (s *DashboardService) DeleteDashboad(id string) (*Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("DELETE", path, nil)
	if rerr != nil {
		return nil, rerr
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
