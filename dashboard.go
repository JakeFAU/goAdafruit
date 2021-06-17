package goadafruit

import (
	"fmt"
	"time"
)

type Dashboard struct {
	Username string `json:"username"`
	Owner    struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"owner"`
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Key          string    `json:"key"`
	Description  string    `json:"description"`
	Visibility   string    `json:"visibility"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ShowHeader   bool      `json:"show_header"`
	ColorMode    string    `json:"color_mode"`
	BlockBorders bool      `json:"block_borders"`
	Layouts      struct {
		Xl []struct {
			X int    `json:"x"`
			Y int    `json:"y"`
			W int    `json:"w"`
			H int    `json:"h"`
			I string `json:"i"`
		} `json:"xl"`
		Lg []struct {
			X int    `json:"x"`
			Y int    `json:"y"`
			W int    `json:"w"`
			H int    `json:"h"`
			I string `json:"i"`
		} `json:"lg"`
		Md []struct {
			X int    `json:"x"`
			Y int    `json:"y"`
			W int    `json:"w"`
			H int    `json:"h"`
			I string `json:"i"`
		} `json:"md"`
		Sm []struct {
			X int    `json:"x"`
			Y int    `json:"y"`
			W int    `json:"w"`
			H int    `json:"h"`
			I string `json:"i"`
		} `json:"sm"`
		Xs []struct {
			X int    `json:"x"`
			Y int    `json:"y"`
			W int    `json:"w"`
			H int    `json:"h"`
			I string `json:"i"`
		} `json:"xs"`
	} `json:"layouts"`
	HeaderImageURL interface{} `json:"header_image_url"`
	Blocks         []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		VisualType      string `json:"visual_type"`
		BlockProperties struct {
			XAxisLabel    string `json:"xAxisLabel"`
			YAxisLabel    string `json:"yAxisLabel"`
			YAxisMin      string `json:"yAxisMin"`
			YAxisMax      string `json:"yAxisMax"`
			DecimalPlaces string `json:"decimalPlaces"`
			RawDataOnly   bool   `json:"rawDataOnly"`
			SteppedLine   bool   `json:"steppedLine"`
			FeedKeyLegend bool   `json:"feedKeyLegend"`
			GridLines     bool   `json:"gridLines"`
			HistoryHours  string `json:"historyHours"`
		} `json:"properties,omitempty"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Source     interface{} `json:"source"`
		SourceKey  interface{} `json:"source_key"`
		BlockFeeds []struct {
			ID   int `json:"id"`
			Feed struct {
				Owner struct {
					Username string `json:"username"`
					ID       int    `json:"id"`
				} `json:"owner"`
				ID         int         `json:"id"`
				Key        string      `json:"key"`
				Name       string      `json:"name"`
				UserID     int         `json:"user_id"`
				Username   string      `json:"username"`
				UpdatedAt  time.Time   `json:"updated_at"`
				UnitType   interface{} `json:"unit_type"`
				UnitSymbol interface{} `json:"unit_symbol"`
				History    bool        `json:"history"`
				Enabled    bool        `json:"enabled"`
				IsShared   bool        `json:"is_shared"`
				Writable   bool        `json:"writable"`
				Visibility string      `json:"visibility"`
			} `json:"feed"`
			Group struct {
				Owner struct {
					ID int `json:"id"`
				} `json:"owner"`
				ID       int    `json:"id"`
				Key      string `json:"key"`
				Name     string `json:"name"`
				UserID   int    `json:"user_id"`
				IsShared bool   `json:"is_shared"`
				Writable bool   `json:"writable"`
			} `json:"group"`
		} `json:"block_feeds"`
		Properties struct {
			ShowIcon      bool   `json:"showIcon"`
			Label         string `json:"label"`
			MinValue      string `json:"minValue"`
			MaxValue      string `json:"maxValue"`
			RingWidth     string `json:"ringWidth"`
			MinWarning    string `json:"minWarning"`
			MaxWarning    string `json:"maxWarning"`
			DecimalPlaces string `json:"decimalPlaces"`
		} `json:"properties,omitempty"`
	} `json:"blocks"`
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

	return d, resp, nil

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
