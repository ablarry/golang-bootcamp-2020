package api

import (
	"github.com/ablarry/golang-bootcamp-2020/pkg/pokemons/controller"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/rest"
	"github.com/gorilla/mux"
	"net/http"
)

const PATH string = "/pokemons"

func Initialize(r *mux.Router) *mux.Router {
	c := &controller.Pokemon{
		ParseURL: rest.DecodeURL,
	}

	api := r.PathPrefix(PATH).Subrouter()
	api.HandleFunc("", c.GetList).Methods(http.MethodGet)
	api.HandleFunc("/{id}", c.GetOne).Methods(http.MethodGet)
	return api
}