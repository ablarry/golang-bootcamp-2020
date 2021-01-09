package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TestDelete test Delete mocking service in controller
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	table := []struct {
		name string
		*http.Request
		code        int
		mockService func(c *gomock.Controller) *service.MockService
		params      map[string]string
		expected    int64
		ok          bool
	}{
		{
			"1._ OK: Delete",
			newRequest(http.MethodDelete, PokemonEndpoint, ""),
			http.StatusOK,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().Delete(gomock.Any()).Return(int64(1), nil).AnyTimes()
				return m
			},
			map[string]string{
				"id": "1",
			},
			1,
			true,
		},
		{
			"2._ Not Content: Delete ",
			newRequest(http.MethodDelete, PokemonEndpoint, ""),
			http.StatusNoContent,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				m.EXPECT().Delete(gomock.Any()).Return(int64(0), nil).AnyTimes()
				return m
			},
			map[string]string{
				"id": "1",
			},
			0,
			true,
		},
		{
			"3._ Not params: Delete ",
			newRequest(http.MethodDelete, PokemonEndpoint, ""),
			http.StatusBadRequest,
			func(c *gomock.Controller) *service.MockService {
				m := service.NewMockService(ctrl)
				return m
			},
			map[string]string{},
			0,
			false,
		},
	}
	for _, tt := range table {
		c := &Pokemon{tt.mockService(ctrl)}
		w := httptest.NewRecorder()
		h := http.HandlerFunc(c.Delete)
		tt.Request = mux.SetURLVars(tt.Request, tt.params)
		h.ServeHTTP(w, tt.Request)

		assert.EqualValues(t, tt.code, w.Code, tt.name)
	}
	ctrl.Finish()
}
