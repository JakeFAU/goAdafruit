package goadafruit

import "fmt"

type Block struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Key         string       `json:"key"`
	VisualType  string       `json:"visual_type"`
	Column      int          `json:"column"`
	Row         int          `json:"row"`
	SizeX       int          `json:"size_x"`
	SizeY       int          `json:"size_y"`
	BlockFeeds  []BlockFeeds `json:"block_feeds"`
}
type BlockFeeds struct {
	ID    string `json:"id"`
	Feed  string `json:"feed"`
	Group string `json:"group"`
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
