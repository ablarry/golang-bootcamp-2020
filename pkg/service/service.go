package service

import (
	"github.com/ablarry/golang-bootcamp-2020/pkg/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"
)

// Service layer business logic
type Service interface {
	GetOne(id int64) (*model.Pokemon, error)
	GetList(paginator, size int64) ([]*model.Pokemon, error)
	Create(*model.Pokemon) (*model.Pokemon, error)
	Delete(id int64) (int64, error)
}

// PokemonService implements Service
type PokemonService struct {
	Repository repo.Repo
}

// GetOne get one element by id
func (s PokemonService) GetOne(id int64) (*model.Pokemon, error) {
	p, err := s.Repository.FindOne(id)
	if err != nil {
		return nil, Handler(err)
	}
	return p, nil
}

// GetList get a list of registries
func (s PokemonService) GetList(paginator, size int64) ([]*model.Pokemon, error) {
	p, err := s.Repository.Find(paginator, size)
	if err != nil {
		return nil, Handler(err)
	}
	return p, nil
}

// Create creates a registry
func (s PokemonService) Create(p *model.Pokemon) (*model.Pokemon, error) {
	pn, err := s.Repository.Create(p)
	if err != nil {
		return nil, Handler(err)
	}
	return pn, nil
}

// Delete deletes a registry
func (s PokemonService) Delete(id int64) (int64, error) {
	rows, err := s.Repository.Delete(id)
	if err != nil {
		return 0, Handler(err)
	}
	return rows, nil
}
