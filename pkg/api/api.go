package api

import (
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/controller"
	"github.com/ablarry/golang-bootcamp-2020/pkg/repo"
	"github.com/ablarry/golang-bootcamp-2020/pkg/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Initialize initializes controller and endpoint
func Initialize(r *mux.Router, db *sqlx.DB) *mux.Router {
	c := &controller.Pokemon{
		S: service.PokemonService{Repository: &repo.PokemonRepo{DB: db}},
	}
	api := r.PathPrefix("/pokemons").Subrouter()
	api.HandleFunc("", c.GetList).Methods(http.MethodGet)
	api.HandleFunc("/{id}", c.GetOne).Methods(http.MethodGet)
	api.HandleFunc("/{id}", c.Delete).Methods(http.MethodDelete)
	api.HandleFunc("", c.Create).Methods(http.MethodPost)
	return api
}
