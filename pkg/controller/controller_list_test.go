package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"
	"github.com/ablarry/golang-bootcamp-2020/pkg/service"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const PokemonsEndpoint = "/v1/pokemons"

// TestList test GetList  mocking service in controller
func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	table := []struct {
		name string
		*http.Request
		code        int
		mockService func(c *gomock.Controller) *service.MockService
		params      map[string]string
		expected    []*model.Pokemon
		ok          bool
	}{
		{
			"1._ OK: GetList ",
			newRequest(http.MethodGet, PokemonsEndpoint, ""),
			http.StatusOK,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(repo.PMocks, nil).AnyTimes()
				return m
			},
			map[string]string{
				"size":      "10",
				"Paginator": "0",
			},
			repo.PMocks,
			true,
		},
		{
			"2._ Empty: GetList",
			newRequest(http.MethodGet, PokemonsEndpoint, ""),
			http.StatusOK,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]*model.Pokemon{}, nil).AnyTimes()
				return m
			},
			map[string]string{
				"size":      "10",
				"Paginator": "0",
			},
			[]*model.Pokemon{},
			true,
		},
		{
			"3._ InvalidQueryParam: GetList",
			newRequest(http.MethodGet, PokemonsEndpoint, ""),
			http.StatusBadRequest,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]*model.Pokemon{}, nil).AnyTimes()
				return m
			},
			map[string]string{
				"size":      "e",
				"paginator": "e",
			},
			[]*model.Pokemon{},
			false,
		},
		{
			"4._ NotQueryParam: GetList",
			newRequest(http.MethodGet, PokemonsEndpoint, ""),
			http.StatusOK,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]*model.Pokemon{}, nil).AnyTimes()
				return m
			},
			map[string]string{},
			[]*model.Pokemon{},
			true,
		},
	}
	for _, tt := range table {
		c := &Pokemon{tt.mockService(ctrl)}
		w := httptest.NewRecorder()
		h := http.HandlerFunc(c.GetList)
		tt.Request = mux.SetURLVars(tt.Request, tt.params)
		h.ServeHTTP(w, tt.Request)

		assert.EqualValues(t, tt.code, w.Code, tt.name)
		if tt.ok {
			expected, err := json.Marshal(tt.expected)
			assert.NoError(t, err, tt.name)
			assert.Equal(t, string(expected), w.Body.String(), tt.name)
		}
	}
	ctrl.Finish()
}
