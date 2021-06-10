package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriggersAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/triggers",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
		[
		{
			"name": "trigger1"
		},
		{
			"name": "trigger2"
		}		
		]`)
		})

	assert := assert.New(t)

	triggers, response, err := client.Trigger.All()

	assert.Nil(err)
	assert.NotNil(triggers)
	assert.NotNil(response)

	assert.Equal(2, len(triggers))
}

func TestCreateTrigger(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/triggers",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `
			{
				"name": "trigger1"
			}`)
		})

	assert := assert.New(t)
	trig := Trigger{Name: "trigger1"}

	trigger, response, err := client.Trigger.Create(&trig)

	assert.Nil(err)
	assert.NotNil(trigger)
	assert.NotNil(response)

	assert.Equal("trigger1", trigger.Name)

}

func TestGetTrigger(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/triggers/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
			{
				"name": "trigger1"
			}`)
		})

	assert := assert.New(t)
	id := "1"

	trigger, response, err := client.Trigger.GetTrigger(&id)

	assert.Nil(err)
	assert.NotNil(trigger)
	assert.NotNil(response)

	assert.Equal("trigger1", trigger.Name)

}

func TestReplaceTrigger(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/triggers/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			fmt.Fprint(w, `
			{
				"name": "trigger2"
			}`)
		})

	assert := assert.New(t)
	id := "1"
	trig := Trigger{Name: "trigger2"}

	trigger, response, err := client.Trigger.ReplaceTrigger(&id, &trig)
	assert.Nil(err)
	assert.NotNil(trigger)
	assert.NotNil(response)

	assert.Equal("trigger2", trigger.Name)
}

func TestDeleteTrigger(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/triggers/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			fmt.Fprint(w, `{"success" : "delete"}`)
		})

	assert := assert.New(t)
	id := "1"
	trig := Trigger{Name: "trigger1"}

	response, err := client.Trigger.DeleteTrigger(&id, &trig)
	assert.Nil(err)
	assert.NotNil(response)

}
