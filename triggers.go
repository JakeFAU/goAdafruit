package goadafruit

import "fmt"

type Trigger struct {
	Name string `json:"name"`
}

type TriggerService struct {
	client *Client
}

func (s *TriggerService) All() ([]*Trigger, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/triggers", s.client.Username)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates Triggers slice
	triggers := make([]*Trigger, 0)
	resp, err := s.client.Do(req, &triggers)
	if err != nil {
		return nil, resp, err
	}

	return triggers, resp, nil
}

func (s *TriggerService) Create(t *Trigger) (*Trigger, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/triggers", s.client.Username)

	req, rerr := s.client.NewRequest("POST", path, t)
	if rerr != nil {
		return nil, nil, rerr
	}

	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

func (s *TriggerService) GetTrigger(id *string) (*Trigger, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/triggers/%v", s.client.Username, *id)

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}
	var trigger Trigger

	resp, err := s.client.Do(req, &trigger)
	if err != nil {
		return nil, resp, err
	}

	return &trigger, resp, nil

}

func (s TriggerService) ReplaceTrigger(id *string, t *Trigger) (*Trigger, *Response, error) {
	path := fmt.Sprintf("api/v2/%v/triggers/%v", s.client.Username, *id)

	req, rerr := s.client.NewRequest("PUT", path, t)
	if rerr != nil {
		return nil, nil, rerr
	}

	var trigger Trigger

	resp, err := s.client.Do(req, &trigger)
	if err != nil {
		return nil, resp, err
	}

	return &trigger, resp, nil

}

func (s TriggerService) DeleteTrigger(id *string, t *Trigger) (*Response, error) {
	path := fmt.Sprintf("api/v2/%v/triggers/%v", s.client.Username, *id)

	req, rerr := s.client.NewRequest("DELETE", path, t)
	if rerr != nil {
		return nil, rerr
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
