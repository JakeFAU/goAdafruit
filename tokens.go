package goadafruit

import "fmt"

type Token struct {
	Token string `json:"token"`
}

type TokenService struct {
	client *Client
}

func (s *TokenService) GetAllTokens() ([]*Token, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	tokens := make([]*Token, 0)

	resp, err := s.client.Do(req, &tokens)
	if err != nil {
		return nil, resp, err
	}

	return tokens, resp, nil

}

func (s *TokenService) CreateToken(t Token) (*Token, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, t)
	if rerr != nil {
		return nil, nil, rerr
	}

	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return &t, resp, nil
}

func (s *TokenService) GetToken(id string) (*Token, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var t Token

	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return &t, resp, nil
}

func (s *TokenService) DeleteToken(id string, tok Token) (*Response, error) {
	path := fmt.Sprintf("/api/v2/%v/tokens/%v", s.client.Username, id)

	req, rerr := s.client.NewRequest("DELETE", path, tok)
	if rerr != nil {
		return nil, rerr
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
