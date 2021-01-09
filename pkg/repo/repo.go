package repo

import (
	"github.com/ablarry/golang-bootcamp-2020/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Queries sql constants
const (
	pokemonFields           = `id, name, type, category, weakness`
	pokemonSelectBase       = `SELECT * FROM pokemon`
	QueryPokemonSelectByID  = pokemonSelectBase + ` WHERE id=$1`
	QueryPokemonSelectLimit = pokemonSelectBase + ` OFFSET $1 LIMIT $2`
	InsertPokemon           = ` INSERT INTO pokemon (` + pokemonFields + `) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	DeletePokemon           = ` DELETE FROM pokemon WHERE id=$1`
)

// Repo int64erface CRUD Pokemon
type Repo interface {
	FindOne(id int64) (*model.Pokemon, error)
	Find(paginator, limit int64) ([]*model.Pokemon, error)
	Create(p *model.Pokemon) (*model.Pokemon, error)
	Delete(id int64) (int64, error)
}

// PokemonRepo implments int64erface Repo
type PokemonRepo struct {
	DB *sqlx.DB
}

// FindOne finds a specific register
func (r PokemonRepo) FindOne(id int64) (*model.Pokemon, error) {
	p := &model.Pokemon{}
	err := r.DB.Get(p, QueryPokemonSelectByID, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Find finds registers with limit size and paginator
func (r PokemonRepo) Find(paginator, size int64) ([]*model.Pokemon, error) {
	p := []*model.Pokemon{}
	err := r.DB.Select(&p, QueryPokemonSelectLimit, paginator, size)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Create creates a new register
func (r PokemonRepo) Create(p *model.Pokemon) (*model.Pokemon, error) {
	stmt, err := r.DB.Prepare(InsertPokemon)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id int64
	err = stmt.QueryRow(p.ID, p.Name, p.Type, p.Category, p.Weakness).Scan(&id)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

// Delete delete a register
func (r PokemonRepo) Delete(id int64) (int64, error) {
	res, err := r.DB.Exec(DeletePokemon, id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
