package goadafruit

import (
	"fmt"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	ID         interface{} `json:"id"`
	Key        string      `json:"key"`
	Master     bool        `json:"master"`
	Createable bool        `json:"createable"`
	Readable   bool        `json:"readable"`
	Updateable bool        `json:"updateable"`
	Deleteable bool        `json:"deleteable"`
	Expiration interface{} `json:"expiration"`
	UserID     int         `json:"user_id"`
	FeedID     interface{} `json:"feed_id"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	QrCode     string      `json:"qr_code"`
	Status     string      `json:"status"`
}

type TokenService struct {
	client *Client
}

func (s *TokenService) GetAllTokens() ([]*TokenResponse, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	tokens := make([]*TokenResponse, 0)

	resp, err := s.client.Do(req, &tokens)
	if err != nil {
		return nil, resp, err
	}

	return tokens, resp, nil

}

func (s *TokenService) CreateToken(t Token) (*TokenResponse, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, t)
	if rerr != nil {
		return nil, nil, rerr
	}

	tr := TokenResponse{}
	resp, err := s.client.Do(req, &tr)
	if err != nil {
		return nil, resp, err
	}

	return &tr, resp, nil
}

func (s *TokenService) GetToken(id string) (*TokenResponse, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var t TokenResponse

	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return &t, resp, nil
}

func (s *TokenService) DeleteToken(id string) (*Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens/%v", s.client.Username, id)

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
