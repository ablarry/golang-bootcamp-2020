package service

import (
	"database/sql"
	"testing"

	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGetList test get list of registers
func TestGetList(t *testing.T) {

	ctrl := gomock.NewController(t)
	table := []struct {
		name     string
		mockRepo func(c *gomock.Controller) *repo.MockRepo
		ok       bool
	}{
		{
			"1._ Retrieve registers ",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(repo.PMocks, nil).AnyTimes()
				return m
			},
			true,
		}, {
			"2._ Retrieve registers with error - sqlError",
			func(c *gomock.Controller) *repo.MockRepo {
				m := repo.NewMockRepo(ctrl)
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, sql.ErrConnDone).AnyTimes()
				return m
			},
			false,
		},
	}
	for _, tt := range table {
		r := tt.mockRepo(ctrl)
		s := &PokemonService{r}
		p, err := s.GetList(0, 10)
		assert.Equal(t, tt.ok, err == nil, "Error not expected")
		if tt.ok {
			assert.Equal(t, p, repo.PMocks, "Not equal registies")

		}

	}
	ctrl.Finish()
}
