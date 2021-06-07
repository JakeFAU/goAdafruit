package goadafruit

import "fmt"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Username  string `json:"username"`
	TimeZone  string `json:"time_zone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserDetails struct {
	DataRateLimit  int `json:"data_rate_limit"`
	ActiveDataRate int `json:"active_data_rate"`
}

type UserService struct {
	client *Client
}

func (s *UserService) GetUserInformation() (*User, *Response, error) {
	path := fmt.Sprint("/api/v2/user")

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

func (u *UserService) GetDetailedUserInformation() (*UserDetails, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/throttle", u.client.Username)

	req, rerr := u.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var user UserDetails
	resp, err := u.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
