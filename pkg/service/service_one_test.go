package service

import (
	"database/sql"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGetOne test get one register
func TestGetOne(t *testing.T) {

	ctrl := gomock.NewController(t)
	table := []struct {
		name     string
		mockRepo func(c *gomock.Controller) *repo.MockRepo
		ok       bool
	}{
		{
			"1._ Retrieve register ",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().FindOne(gomock.Any()).Return(repo.PMock, nil).AnyTimes()
				return m
			},
			true,
		}, {
			"2._ Retrieve register with error - sql.ErrNoRows ",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().FindOne(gomock.Any()).Return(nil, sql.ErrNoRows).AnyTimes()
				return m
			},
			false,
		},
	}
	for _, tt := range table {
		r := tt.mockRepo(ctrl)
		s := &PokemonService{r}
		p, err := s.GetOne(1)
		assert.Equal(t, tt.ok, err == nil, "Error not expected")
		if tt.ok {
			assert.Equal(t, p, repo.PMock, "Not equal pokemons registry")
		}

	}
	ctrl.Finish()
}
