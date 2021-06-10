package goadafruit

import "fmt"

type Dashboard struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Key         string  `json:"key"`
	Blocks      []Block `json:"blocks"`
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

func (s *DashboardService) CreateDashboard(d Dashboard) (*Dashboard, *Response, error) {
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
