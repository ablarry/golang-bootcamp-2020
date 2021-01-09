package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"
	"github.com/ablarry/golang-bootcamp-2020/pkg/service"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const PokemonEndpoint = "/v1/pokemons/1"

// TestOne test GetOne mocking service in controller
func TestOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	table := []struct {
		name string
		*http.Request
		code        int
		mockService func(c *gomock.Controller) *service.MockService
		params      map[string]string
		ok          bool
	}{
		{
			"1._ OK: GetOne ",
			newRequest(http.MethodGet, PokemonEndpoint, ""),
			http.StatusOK,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetOne(gomock.Any()).Return(repo.PMock, nil).AnyTimes()
				return m
			},
			map[string]string{
				"id": "1",
			},
			true,
		},
		{
			"2._ NotFound: GetOne ",
			newRequest(http.MethodGet, PokemonEndpoint, ""),
			http.StatusNotFound,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetOne(gomock.Any()).Return(nil, errors.ErrNotFound).AnyTimes()
				return m
			},
			map[string]string{
				"id": "1",
			},
			false,
		},
		{
			"3._ Not params : GetOne ",
			newRequest(http.MethodGet, PokemonEndpoint, ""),
			http.StatusBadRequest,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				return m
			},
			map[string]string{},
			false,
		},
	}
	for _, tt := range table {
		c := &Pokemon{tt.mockService(ctrl)}
		w := httptest.NewRecorder()
		h := http.HandlerFunc(c.GetOne)
		tt.Request = mux.SetURLVars(tt.Request, tt.params)
		h.ServeHTTP(w, tt.Request)

		assert.EqualValues(t, tt.code, w.Code, tt.name)
		if tt.ok {
			expected, err := json.Marshal(repo.PMock)
			assert.NoError(t, err, tt.name)
			assert.Equal(t, string(expected), w.Body.String(), tt.name)
		}
	}
	ctrl.Finish()
}

// newRequest helper to create request
func newRequest(method, uri, body string) *http.Request {
	req, _ := http.NewRequest(method, uri, bytes.NewBufferString(body))
	return req
}
