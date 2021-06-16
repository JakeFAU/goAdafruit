package goadafruit

import (
	"fmt"
	"time"
)

type Block struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	VisualType string       `json:"visual_type"`
	Column     interface{}  `json:"column"`
	Row        interface{}  `json:"row"`
	SizeX      interface{}  `json:"size_x"`
	SizeY      interface{}  `json:"size_y"`
	SourceKey  interface{}  `json:"source_key"`
	Source     interface{}  `json:"source"`
	Properties Properties   `json:"properties"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	BlockFeeds []BlockFeeds `json:"block_feeds"`
}
type Properties struct {
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
}

type BlockFeeds struct {
	ID    int   `json:"id"`
	Feed  Feed  `json:"feed"`
	Group Group `json:"group"`
}

type BlockService struct {
	client *Client
}

func (s *BlockService) AllBlocks(dashboardID string) ([]*Block, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v/blocks", s.client.Username, dashboardID)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	blocks := make([]*Block, 0)
	resp, err := s.client.Do(req, &blocks)
	if err != nil {
		return nil, resp, err
	}

	return blocks, resp, nil
}

func (s *BlockService) CreateBlock(dashboardID string, b Block) (*Block, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v/blocks", s.client.Username, dashboardID)
	req, rerr := s.client.NewRequest("POST", path, b)
	if rerr != nil {
		return nil, nil, rerr
	}
	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}

	return &b, resp, nil
}

func (s *BlockService) GetBlock(dashboardID string, blockID string) (*Block, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v/blocks/%v", s.client.Username, dashboardID, blockID)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var block Block
	resp, err := s.client.Do(req, &block)
	if err != nil {
		return nil, resp, err
	}

	return &block, resp, nil
}

func (s *BlockService) ReplaceBlock(dashboardID string, blockID string, b Block) (*Block, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v/blocks/%v", s.client.Username, dashboardID, blockID)

	req, rerr := s.client.NewRequest("PUT", path, b)
	if rerr != nil {
		return nil, nil, rerr
	}

	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return &b, resp, nil
}

func (s *BlockService) DeleteBlock(dashboardID string, blockID string) (*Response, error) {
	path := fmt.Sprintf("/api/v2/%v/dashboards/%v/blocks/%v", s.client.Username, dashboardID, blockID)

	req, rerr := s.client.NewRequest("DELETE", path, nil)
	if rerr != nil {
		return nil, rerr
	}

	var block Block
	resp, err := s.client.Do(req, &block)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
