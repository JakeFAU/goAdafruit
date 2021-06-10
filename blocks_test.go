package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBlocks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1/blocks",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id": 0},{"id": 1}]`)
		})

	assert := assert.New(t)

	blocks, response, err := client.Blocks.AllBlocks("1")

	assert.Nil(err)
	assert.NotNil(blocks)
	assert.NotNil(response)

	assert.Equal(2, len(blocks))
}

func TestGetBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1/blocks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id": 1}`)
		})

	assert := assert.New(t)

	block, response, err := client.Blocks.GetBlock("1", "1")

	assert.Nil(err)
	assert.NotNil(block)
	assert.NotNil(response)

}

func TestCreateBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1/blocks",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
		})

	assert := assert.New(t)
	b := Block{Name: "Indoor"}

	block, response, err := client.Blocks.CreateBlock("1", b)

	assert.Nil(err)
	assert.NotNil(block)
	assert.NotNil(response)

}

func TestReplaceBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1/blocks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
		})

	assert := assert.New(t)
	b := Block{Name: "Indoor"}

	block, response, err := client.Blocks.ReplaceBlock("1", "1", b)

	assert.Nil(err)
	assert.NotNil(block)
	assert.NotNil(response)

}

func TestDeleteBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/dashboards/1/blocks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		})

	assert := assert.New(t)
	response, err := client.Blocks.DeleteBlock("1", "1")

	assert.Nil(err)
	assert.NotNil(response)

}
