package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllActivities(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/activities",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id": 0},{"id": 1}]`)
		})

	assert := assert.New(t)

	activities, response, err := client.Activities.All()

	assert.Nil(err)
	assert.NotNil(activities)
	assert.NotNil(response)

	assert.Equal(2, len(activities))

}

func TestGroupActivities(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/activities/move_block",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id": 1}]`)
		})

	assert := assert.New(t)
	group := "move_block"

	activities, response, err := client.Activities.ActivitiesByType(&group)

	assert.Nil(err)
	assert.NotNil(activities)
	assert.NotNil(response)

	assert.Equal(1, len(activities))

}
