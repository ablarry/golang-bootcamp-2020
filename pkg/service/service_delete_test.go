package service

import (
	"database/sql"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestDelete deletes one register
func TestDelete(t *testing.T) {

	ctrl := gomock.NewController(t)
	table := []struct {
		name     string
		mockRepo func(c *gomock.Controller) *repo.MockRepo
		ok       bool
	}{
		{
			"1._ Delete a register ",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().Delete(gomock.Any()).Return(int64(1), nil).AnyTimes()
				return m
			},
			true,
		}, {
			"2._ Delete a register with error ",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().Delete(gomock.Any()).Return(int64(0), sql.ErrConnDone).AnyTimes()
				return m
			},
			false,
		},
	}
	for _, tt := range table {
		r := tt.mockRepo(ctrl)
		s := &PokemonService{r}
		p, err := s.Delete(1)
		assert.Equal(t, tt.ok, err == nil, "Error not expected")
		if tt.ok {
			assert.Equal(t, p, int64(1), "Not equal pokemons registry")
		}

	}
	ctrl.Finish()
}
