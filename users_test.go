package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInformation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/user",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
			{
				"id": 1,
				"name": "Jake",
				"color": "blue",
				"username": "JakeFau",
				"time_zone": "EST"
			}`)
		})

	assert := assert.New(t)

	user, response, err := client.User.GetUserInformation()

	assert.Nil(err)
	assert.NotNil(user)
	assert.NotNil(response)

	assert.Equal("Jake", user.Name)

}

func TestDetailedUserInformation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/throttle",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
		{
			"data_rate_limit": 60,
			"active_data_rate": 20
		}`)
		})

	assert := assert.New(t)

	user, response, err := client.User.GetDetailedUserInformation()

	assert.Nil(err)
	assert.NotNil(user)
	assert.NotNil(response)

	assert.Equal(60, user.DataRateLimit)

}
