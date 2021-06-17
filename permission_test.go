package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPermissions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feed/1/acl",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"ID": 0,"UserID": "jakefau"},{"ID": 1,"UserID": "jakefau"}]`)
		})
	assert := assert.New(t)

	permission, response, err := client.Permission.AllPermissions()

	assert.Nil(err)
	assert.NotNil(permission)
	assert.NotNil(response)

	assert.Equal(2, len(permission))
}

func TestCreatePermission(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feed/1/acl",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
		})
	assert := assert.New(t)

	p := Permission{ID: 1, UserID: 11023}
	permission, response, err := client.Permission.CreatePermission(&p, "feed", "1")

	assert.Nil(err)
	assert.NotNil(permission)
	assert.NotNil(response)

}

func TestGetPermission(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feed/1/acl/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
		})
	assert := assert.New(t)

	permission, response, err := client.Permission.GetPermission("feed", "1", "1")
	assert.Nil(err)
	assert.NotNil(permission)
	assert.NotNil(response)

}

func TestReplacePermission(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feed/1/acl/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
		})
	assert := assert.New(t)
	p := Permission{ID: 1, UserID: 2}

	permission, response, err := client.Permission.ReplacePermission(p, "feed", "1", "1")
	assert.Nil(err)
	assert.NotNil(permission)
	assert.NotNil(response)
	assert.Equal(2, permission.UserID)

}

func TestDeletePermission(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/feed/1/acl/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		})
	assert := assert.New(t)

	response, err := client.Permission.DeletePermission("feed", "1", "1")
	assert.Nil(err)
	assert.NotNil(response)
}
