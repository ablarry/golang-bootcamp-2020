package rest

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type ParamsOne struct {
	ID int64 `form:":id" validate:"required,gt=0"`
}

// TestDecodeURL test mappper params from URL
func TestDecodeURL(t *testing.T) {
	table := []struct {
		name string
		*http.Request
		params map[string]string
		p      *ParamsOne
		ok     bool
	}{
		{
			"1._ OK: Get Params ",
			newRequest("GET", "/v1/pokemons/4", ""),
			map[string]string{
				"id": "4",
			},
			&ParamsOne{4},
			true,
		},
	}

	for _, tt := range table {
		param := &ParamsOne{}
		tt.Request = mux.SetURLVars(tt.Request, tt.params)
		err := DecodeURL(tt.Request, param)
		if tt.ok {
			assert.Equal(t, tt.p, param, "Error param not equal as expected")
		} else {
			assert.Error(t, err, "Expected error")
		}

	}
}

// newRequest helper to create request
func newRequest(method, uri, body string) *http.Request {
	req, _ := http.NewRequest(method, uri, bytes.NewBufferString(body))
	return req
}
