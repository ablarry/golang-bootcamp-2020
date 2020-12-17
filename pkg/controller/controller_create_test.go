package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"
	"github.com/ablarry/golang-bootcamp-2020/pkg/service"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestCreate test Create mocking service in controller
func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	table := []struct {
		name string
		*http.Request
		code        int
		mockService func(c *gomock.Controller) *service.MockService
		params      map[string]string
		expected    *model.Pokemon
		ok          bool
	}{
		{
			"1._ OK: Create",
			newRequest(http.MethodPost, "/v1/pokemons", `{ "id": 1,"category": "Seed",  "name": "Bulbasaur", "type": "Grass", "wakness": "Fire" }`),
			http.StatusCreated,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().Create(gomock.Any()).Return(repo.PMock, nil).AnyTimes()
				return m
			},
			map[string]string{},
			repo.PMock,
			true,
		},
		{
			"2._ Register already exists: Create",
			newRequest(http.MethodPost, "/v1/pokemons", `{ "id": 1,"category": "Seed",  "name": "Bulbasaur", "type": "Grass", "wakness": "Fire" }`),
			http.StatusUnprocessableEntity,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().Create(gomock.Any()).Return(nil, errors.ErrDuplicate).AnyTimes()
				return m
			},
			map[string]string{},
			repo.PMock,
			false,
		},
		{
			"3._ Invalid format : Create",
			newRequest(http.MethodPost, "/v1/pokemons", ` "id": 1, "wakness": "Fire" `),
			http.StatusBadRequest,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				return m
			},
			map[string]string{},
			nil,
			false,
		},
	}
	for _, tt := range table {
		c := &Pokemon{tt.mockService(ctrl)}
		w := httptest.NewRecorder()
		h := http.HandlerFunc(c.Create)
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
