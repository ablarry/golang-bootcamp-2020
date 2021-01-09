package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"gopkg.in/stretchr/testify.v1/assert"
)

var PokemonFields = []string{"id", "name", "type", "category", "weakness"}

const (
	mockPokemonSelectBase       = `SELECT (.+) FROM pokemon`
	MockQueryPokemonSelectByID  = mockPokemonSelectBase + ` WHERE id=?`
	MockQueryPokemonSelectLimit = mockPokemonSelectBase + ` OFFSET \$1 LIMIT \$2`
	MockInsertPokemon           = ` INSERT INTO pokemon \(` + pokemonFields + `\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id`
)

// InitMock initlize mock sql
func InitMock() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("Error to initialize sqlmock")
	}
	return sqlx.NewDb(db, "postgres"), mock

}

// TestFindOne finds one register
func TestFindOne(t *testing.T) {
	db, mock := InitMock()
	repo := &PokemonRepo{db}
	rows := sqlmock.NewRows(PokemonFields).AddRow(PMock.ID, PMock.Name, PMock.Type, PMock.Category, PMock.Weakness)
	mock.ExpectQuery(MockQueryPokemonSelectByID).WithArgs(PMock.ID).WillReturnRows(rows)
	pokemon, err := repo.FindOne(PMock.ID)
	assert.NotNil(t, pokemon)
	assert.NoError(t, err)
}

// TestFind finds registers
func TestFind(t *testing.T) {
	db, mock := InitMock()
	repo := &PokemonRepo{db}
	rows := sqlmock.NewRows(PokemonFields)
	rows.AddRow(PMock.ID, PMock.Name, PMock.Type, PMock.Category, PMock.Weakness)
	rows.AddRow(PMock2.ID, PMock2.Name, PMock2.Type, PMock2.Category, PMock2.Weakness)
	mock.ExpectQuery(MockQueryPokemonSelectLimit).WithArgs(0, 10).WillReturnRows(rows)
	pokemons, err := repo.Find(0, 10)
	assert.NotNil(t, pokemons)
	assert.NoError(t, err)
}

// TestCreate creates register
func TestCreate(t *testing.T) {
	db, mock := InitMock()
	repo := &PokemonRepo{db}
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1)
	mock.ExpectPrepare(MockInsertPokemon).ExpectQuery().WithArgs(PMock.ID, PMock.Name, PMock.Type, PMock.Category, PMock.Weakness).WillReturnRows(rows)
	p, err := repo.Create(PMock)
	assert.NoError(t, err)
	assert.Equal(t, p, PMock)

}
