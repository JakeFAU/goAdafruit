package goadafruit

import (
	"fmt"
	"io"
	"time"
)

// Data are the values contained by a Feed.
type Data struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	FeedID    int    `json:"feed_id"`
	FeedKey   string `json:"feed_key"`
	CreatedAt string `json:"created_at"`
	Location  struct {
	} `json:"location"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	Ele          float64 `json:"ele"`
	CreatedEpoch int     `json:"created_epoch"`
	Expiration   string  `json:"expiration"`
}

type ChartData struct {
	Feed struct {
		ID   string `json:"id"`
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"feed"`
	Parameters struct {
		StartTime  time.Time `json:"start_time"`
		EndTime    time.Time `json:"end_time"`
		Resolution int       `json:"resolution"`
		Hours      int       `json:"hours"`
		Field      string    `json:"field"`
	} `json:"parameters"`
	Columns []string        `json:"columns"`
	Data    [][]interface{} `json:"data"`
}

type DataFilter struct {
	StartTime string `url:"start_time,omitempty"`
	EndTime   string `url:"end_time,omitempty"`
}

type DataService struct {
	client *Client
}

// Create adds a new Data value to an existing Feed.
func (s *DataService) Create(dp *Data) (*Data, *Response, error) {
	path, ferr := s.client.Feed.Path("/data")
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("POST", path, dp)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates a new datapoint
	point := &Data{}
	resp, err := s.client.Do(req, point)
	if err != nil {
		return nil, resp, err
	}

	return point, resp, nil
}

// Create adds new Datam values to an existing Feed.
func (s *DataService) CreateBatch(dp *[]Data) (*Data, *Response, error) {
	path, ferr := s.client.Feed.Path("/data/batch")
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("POST", path, dp)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates a new datapoint
	point := &Data{}
	resp, err := s.client.Do(req, point)
	if err != nil {
		return nil, resp, err
	}

	return point, resp, nil
}

// All returns all Data for the currently selected Feed. See Client.SetFeed()
// for details on selecting a Feed.
func (s *DataService) All(opt *DataFilter) ([]*Data, *Response, error) {
	path, ferr := s.client.Feed.Path("/data")
	if ferr != nil {
		return nil, nil, ferr
	}

	path, oerr := addOptions(path, opt)
	if oerr != nil {
		return nil, nil, oerr
	}

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates Feed slice
	datas := make([]*Data, 0)
	resp, err := s.client.Do(req, &datas)
	if err != nil {
		return nil, resp, err
	}

	return datas, resp, nil
}

// returns the feed data ready for charting
func (s *DataService) GetChartData(opt *DataFilter) (*ChartData, *Response, error) {
	path, ferr := s.client.Feed.Path("/data/chart")
	if ferr != nil {
		return nil, nil, ferr
	}

	path, oerr := addOptions(path, opt)
	if oerr != nil {
		return nil, nil, oerr
	}

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}
	// request populates a new datapoint
	chartData := &ChartData{}
	resp, err := s.client.Do(req, chartData)
	if err != nil {
		return nil, resp, err
	}

	return chartData, resp, nil

}

// Search has the same response format as All, but it accepts optional params
// with which your data can be queried.
func (s *DataService) Search(filter *DataFilter) ([]*Data, *Response, error) {
	path, ferr := s.client.Feed.Path("/data")
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates Feed slice
	datas := make([]*Data, 0)
	resp, err := s.client.Do(req, &datas)
	if err != nil {
		return nil, resp, err
	}

	return datas, resp, nil
}

// Get returns a single Data element, identified by the given ID parameter.
func (s *DataService) Get(id int) (*Data, *Response, error) {
	path, ferr := s.client.Feed.Path(fmt.Sprintf("/data/%v", id))
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var data Data
	resp, err := s.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data, resp, nil
}

// Update takes an ID and a Data record, updates the record idendified by ID,
// and returns a new, updated Data instance.
func (s *DataService) Update(id interface{}, data *Data) (*Data, *Response, error) {
	path, ferr := s.client.Feed.Path(fmt.Sprintf("/data/%v", id))
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("PATCH", path, data)
	if rerr != nil {
		return nil, nil, rerr
	}

	var updatedData Data
	resp, err := s.client.Do(req, &updatedData)
	if err != nil {
		return nil, resp, err
	}

	return &updatedData, resp, nil
}

// Delete the Data identified by the given ID.
func (s *DataService) Delete(id int) (*Response, error) {
	path, ferr := s.client.Feed.Path(fmt.Sprintf("/data/%v", id))
	if ferr != nil {
		return nil, ferr
	}

	req, rerr := s.client.NewRequest("DELETE", path, nil)
	if rerr != nil {
		return nil, rerr
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// private method for handling the Next, Prev, and Last commands
func (s *DataService) retrieve(command string) (*Data, *Response, error) {
	path, ferr := s.client.Feed.Path(fmt.Sprintf("/data/%v", command))
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var data Data
	resp, err := s.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data, resp, nil
}

// Next returns the next Data in the stream.
func (s *DataService) Next() (*Data, *Response, error) {
	return s.retrieve("next")
}

// Prev returns the previous Data in the stream.
func (s *DataService) Prev() (*Data, *Response, error) {
	return s.retrieve("prev")
}

// Last returns the last Data in the stream.
func (s *DataService) Last() (*Data, *Response, error) {
	return s.retrieve("last")
}

// Last returns the first Data in the stream.
func (s *DataService) First() (*Data, *Response, error) {
	return s.retrieve("first")
}

func (s *DataService) MostRecent() (*string, *Response, error) {
	path, ferr := s.client.Feed.Path("/data/retain")
	if ferr != nil {
		return nil, nil, ferr
	}
	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}
	var v interface{}
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	bd := string(b)
	return &bd, resp, nil

}

func (s *DataService) CreateDataInGroup(groupKey string, data *Data) (*Data, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/groups/%v/feeds/%v/data", s.client.Username, groupKey, s.client.Feed.CurrentFeed.ID)

	req, rerr := s.client.NewRequest("POST", path, data)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates a new datapoint
	point := &Data{}
	resp, err := s.client.Do(req, point)
	if err != nil {
		return nil, resp, err
	}

	return point, resp, nil

}

func (s *DataService) CreateDatumInGroup(groupKey string, dp *[]Data) (*Data, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/groups/%v/feeds/%v/data/batch", s.client.Username, groupKey, s.client.Feed.CurrentFeed.ID)

	req, rerr := s.client.NewRequest("POST", path, dp)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates a new datapoint
	point := &Data{}
	resp, err := s.client.Do(req, point)
	if err != nil {
		return nil, resp, err
	}

	return point, resp, nil

}
