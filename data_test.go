package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData_MissingFeed(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{"id":1, "value":"67.112"}`)
		},
	)

	assert := assert.New(t)

	dp := &Data{}
	datapoint, response, err := client.Data.Create(dp)

	assert.NotNil(err)
	assert.Nil(datapoint)
	assert.Nil(response)

	assert.Equal(err.Error(), "CurrentFeed must be set")
}

func TestData_Unauthenticated(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{"id":1, "value":"67.112"}`)
		},
	)

	assert := assert.New(t)

	dp := &Data{}
	datapoint, response, err := client.Data.Create(dp)

	assert.NotNil(err)
	assert.Nil(datapoint)
	assert.Nil(response)

	assert.Equal(err.Error(), "CurrentFeed must be set")
}

func TestDataCreate(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{"id":"1", "value":"67.112"}`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	val := "67.112"

	dp := &Data{
		ID:    "temperature",
		Value: val,
	}
	datapoint, response, err := client.Data.Create(dp)

	assert.Nil(err)
	assert.NotNil(datapoint)
	assert.NotNil(response)

	assert.Equal("1", datapoint.ID)
	assert.Equal(val, datapoint.Value)
}

func TestDataGet(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data/0ERN7NE8FNEK5KGX00SW2GK31B",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"ID":"0ERN7NE8FNEK5KGX00SW2GK31B", "value":"67.112"}`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	datapoint, response, err := client.Data.Get("0ERN7NE8FNEK5KGX00SW2GK31B")

	assert.Nil(err)
	assert.NotNil(datapoint)
	assert.NotNil(response)

	assert.Equal("0ERN7NE8FNEK5KGX00SW2GK31B", datapoint.ID)
	assert.Equal("67.112", datapoint.Value)
}

func TestAllDataNoFilter(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id":"1", "value":"67.112"}]`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	// with no params
	datapoints, response, err := client.Data.All(nil)
	datapoint := datapoints[0]

	assert.Nil(err)
	assert.NotNil(datapoint)
	assert.NotNil(response)

	assert.Equal("1", datapoint.ID)
	assert.Equal("67.112", datapoint.Value)
}

func TestAllDataFilter(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQuery(t, r, "start_time", "2000-01-01")
			testQuery(t, r, "end_time", "2010-01-01")
			fmt.Fprint(w, `[{"id":"1", "value":"67.112"}]`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	// with no params
	datapoints, response, err := client.Data.All(&DataFilter{
		StartTime: "2000-01-01",
		EndTime:   "2010-01-01",
	})
	datapoint := datapoints[0]

	assert.Nil(err)
	assert.NotNil(datapoint)
	assert.NotNil(response)

	assert.Equal("1", datapoint.ID)
	assert.Equal("67.112", datapoint.Value)
}

func TestDataDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feeds/test/data/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "test"})

	response, err := client.Data.Delete("1")

	assert.Nil(err)
	assert.NotNil(response)

	assert.Equal(200, response.StatusCode)
}

func TestDataQueue(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data/next",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id":"1", "value":"1"}`)
		},
	)

	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data/previous",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id":"2", "value":"2"}`)
		},
	)

	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data/last",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id":"3", "value":"3"}`)
		},
	)

	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data/retain",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id":"4", "value":"4"}`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	var (
		datapoint *Data
		response  *Response
		err       error
	)

	datapoint, response, err = client.Data.Next()
	assert.Nil(err)
	assert.NotNil(response)
	assert.Equal("1", datapoint.ID)
	assert.Equal("1", datapoint.Value)

	datapoint, response, err = client.Data.Prev()
	assert.Nil(err)
	assert.NotNil(response)
	assert.Equal("2", datapoint.ID)
	assert.Equal("2", datapoint.Value)

	datapoint, response, err = client.Data.Last()
	assert.Nil(err)
	assert.NotNil(response)
	assert.Equal("3", datapoint.ID)
	assert.Equal("3", datapoint.Value)
}

func TestChartData(t *testing.T) {
	setup()
	defer teardown()

	// prepare endpoint URL for just this request
	mux.HandleFunc("/api/v2/test-user/feeds/temperature/data",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQuery(t, r, "start_time", "2000-01-01")
			testQuery(t, r, "end_time", "2010-01-01")
			fmt.Fprint(w, `[{"id":"1", "value":"67.112"}]`)
		},
	)

	assert := assert.New(t)

	client.SetFeed(&Feed{Key: "temperature"})

	// with no params
	datapoints, response, err := client.Data.All(&DataFilter{
		StartTime: "2000-01-01",
		EndTime:   "2010-01-01",
	})
	datapoint := datapoints[0]

	assert.Nil(err)
	assert.NotNil(datapoint)
	assert.NotNil(response)

	assert.Nil(err)
	assert.NotNil(response)
	assert.NotNil(datapoints)

}
