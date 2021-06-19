package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTokens(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/tokens",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id":"Iamatoken"}, {"id":"Iamanothertoken"}]`)
		},
	)

	assert := assert.New(t)

	tokens, response, err := client.Tokens.GetAllTokens()

	assert.Nil(err)
	assert.NotNil(tokens)
	assert.NotNil(response)

	assert.Equal(2, len(tokens))

}

func TestGetToken(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/tokens/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id":"1"}`)
		},
	)

	assert := assert.New(t)

	token, response, err := client.Tokens.GetToken("1")
	assert.Nil(err)
	assert.NotNil(token)
	assert.NotNil(response)

}

func TestDeleteToken(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/api/v2/test-user/tokens/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			fmt.Fprint(w, "token")
		},
	)

	assert := assert.New(t)

	response, err := client.Tokens.DeleteToken("1")
	assert.Nil(err)
	assert.NotNil(response)
}
