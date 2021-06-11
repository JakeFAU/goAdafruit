package goadafruit

type DateStruct struct {
	Year  int `json:"year"`
	Mon   int `json:"mon"`
	Mday  int `json:"mday"`
	Hour  int `json:"hour"`
	Min   int `json:"min"`
	Sec   int `json:"sec"`
	Wday  int `json:"wday"`
	Yday  int `json:"yday"`
	Isdst int `json:"isdst"`
}

type TimeService struct {
	client *Client
}

func (s *TimeService) UnixSeconds() (*int, *Response, error) {
	path := "/api/v2/time/seconds"

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var tm int
	resp, err := s.client.Do(req, tm)
	if err != nil {
		return &tm, resp, err
	}
	return &tm, resp, nil

}

func (s *TimeService) UnixMilliSeconds() (*float32, *Response, error) {
	path := "/api/v2/time/millis"

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var tm float32
	resp, err := s.client.Do(req, tm)
	if err != nil {
		return &tm, resp, err
	}
	return &tm, resp, nil
}

func (s *TimeService) ISOTime() (*string, *Response, error) {
	path := "/api/v2/time/ISO-8601"

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var tm string
	resp, err := s.client.Do(req, tm)
	if err != nil {
		return &tm, resp, err
	}
	return &tm, resp, nil
}
