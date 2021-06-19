package goadafruit

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// Data are the values contained by a Feed.
type Data struct {
	ID        string `json:"id,omitempty"`
	Value     string `json:"value,omitempty"`
	FeedID    int    `json:"feed_id,omitempty"`
	FeedKey   string `json:"feed_key,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Location  struct {
	} `json:"location,omitempty"`
	Lat          string `json:"lat,omitempty"`
	Lon          string `json:"lon,omitempty"`
	Ele          string `json:"ele,omitempty"`
	CreatedEpoch int    `json:"created_epoch,omitempty"`
	Expiration   string `json:"expiration,omitempty"`
}

type DataPoint struct {
	X string `json:"X"`
	Y string `json:"Y"`
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

// Create adds new Data values to an existing Feed.
func (s *DataService) CreateBatch(dp *[]Data) (*[]Data, *Response, error) {
	path, ferr := s.client.Feed.Path("/data/batch")
	if ferr != nil {
		return nil, nil, ferr
	}

	req, rerr := s.client.NewRequest("POST", path, dp)
	if rerr != nil {
		return nil, nil, rerr
	}

	// request populates a new datapoint
	point := &[]Data{}
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

	for moredata(resp) {
		links := resp.Header.Get("Link")
		dataLink := strings.Split(links, ",")[1]
		dataLink = dataLink[2 : len(dataLink)-1]
		req, _ = s.client.NewRequest("GET", dataLink, nil)
		d := make([]*Data, 0)
		resp, err = s.client.Do(req, &d)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, d...)

	}

	return datas, resp, nil
}

// returns the feed data ready for charting
func (s *DataService) GetChartData(opt *DataFilter) ([]*DataPoint, *Response, error) {
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

	// request populates Feed slice
	var datas map[string]interface{}
	resp, err := s.client.Do(req, &datas)
	if err != nil {
		return nil, resp, err
	}
	var points []*DataPoint
	for k, v := range datas {
		if k == "data" {
			for _, vv := range v.([]interface{}) {
				var point DataPoint
				for _, vvv := range vv.([]interface{}) {
					if point.X == "" {
						point.X = vvv.(string)
					} else {
						point.Y = vvv.(string)
					}
				}
				points = append(points, &point)
			}
		}
	}

	for moredata(resp) {
		links := resp.Header.Get("Link")
		dataLink := strings.Split(links, ",")[1]
		dataLink = dataLink[2 : len(dataLink)-1]
		req, _ = s.client.NewRequest("GET", dataLink, nil)
		resp, err = s.client.Do(req, &datas)
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range datas {
			if k == "data" {
				for _, vv := range v.([]interface{}) {
					var point DataPoint
					for _, vvv := range vv.([]interface{}) {
						if point.X == "" {
							point.X = vvv.(string)
						} else {
							point.Y = vvv.(string)
						}
					}
					points = append(points, &point)
				}
			}
		}

	}

	return points, resp, nil

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
	return s.retrieve("previous")
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

	b, _ := io.ReadAll(resp.Body)
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

func moredata(r *Response) bool {
	count, _ := strconv.ParseInt(r.Header.Get("X-Pagination-Count"), 10, 64)
	limit, _ := strconv.ParseInt(r.Header.Get("X-Pagination-Limit"), 10, 64)
	total, _ := strconv.ParseInt(r.Header.Get("X-Pagination-Total"), 10, 64)

	return (count == limit) && total > limit
}
