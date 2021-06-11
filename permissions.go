package goadafruit

import "fmt"

type Permission struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Scope      string `json:"scope"`
	ScopeValue string `json:"scope_value"`
	Model      string `json:"model"`
	ObjectID   int    `json:"object_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PermissionService struct {
	client *Client
}

func (s *PermissionService) AllPermissions() ([]*Permission, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/%v/%v/acl", s.client.Username, "feed", "1")

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	perms := make([]*Permission, 0)

	resp, err := s.client.Do(req, &perms)
	if err != nil {
		return nil, resp, err
	}

	return perms, resp, nil
}

func (s *PermissionService) CreatePermission(p Permission, ptype string, typeid string) (*Permission, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/%v/%v/acl", s.client.Username, ptype, typeid)

	req, rerr := s.client.NewRequest("POST", path, p)
	if rerr != nil {
		return nil, nil, rerr
	}
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return &p, resp, nil
}

func (s *PermissionService) GetPermission(ptype string, typeid string, pid string) (*Permission, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/%v/%v/acl/%v", s.client.Username, ptype, typeid, pid)
	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}
	var p Permission
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return &p, resp, nil

}

func (s *PermissionService) ReplacePermission(p Permission, ptype string, typeid string, pid string) (*Permission, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/%v/%v/acl/%v", s.client.Username, ptype, typeid, pid)
	req, rerr := s.client.NewRequest("PUT", path, p)
	if rerr != nil {
		return nil, nil, rerr
	}
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return &p, resp, nil
}

func (s *PermissionService) DeletePermission(ptype string, typeid string, pid string) (*Response, error) {
	path := fmt.Sprintf("/api/v2/%v/%v/%v/acl/%v", s.client.Username, ptype, typeid, pid)
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
