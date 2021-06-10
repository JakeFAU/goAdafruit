package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllDashboards(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id": 0},{"id": 1}]`)
		})

	assert := assert.New(t)

	dashboard, response, err := client.Dashboard.AllDashboards()

	assert.Nil(err)
	assert.NotNil(dashboard)
	assert.NotNil(response)

	assert.Equal(2, len(dashboard))

}

func TestCreateDashboard(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
		})

	assert := assert.New(t)

	d := Dashboard{Name: "Dashboard1"}
	dashboard, response, err := client.Dashboard.CreateDashboard(d)

	assert.Nil(err)
	assert.NotNil(dashboard)
	assert.NotNil(response)

	assert.Equal("Dashboard1", dashboard.Name)

}

func TestGetDashboard(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"Name": "DB One"}`)
		})

	assert := assert.New(t)
	dashboard, response, err := client.Dashboard.GetDashboard("1")
	assert.Nil(err)
	assert.NotNil(dashboard)
	assert.NotNil(response)

	assert.Equal("DB One", dashboard.Name)

}

func TestChangeDashboard(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			fmt.Fprint(w, `{"Name": "DB One"}`)
		})

	assert := assert.New(t)
	newdb := Dashboard{Name: "DB1", Description: "Test Dashboard"}
	dashboard, response, err := client.Dashboard.ChangeDashboard("1", newdb)
	assert.Nil(err)
	assert.NotNil(dashboard)
	assert.NotNil(response)

	assert.Equal("DB One", dashboard.Name)

}

func TestDeleteDashboard(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		})

	assert := assert.New(t)
	response, err := client.Dashboard.DeleteDashboad("1")

	assert.Nil(err)
	assert.NotNil(response)

}
