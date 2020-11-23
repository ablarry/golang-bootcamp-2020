package controller

import (
	"fmt"
	"github.com/ablarry/golang-bootcamp-2020/pkg/pokemons/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/render"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/rest"
	"net/http"
)

// Pokemon is a specific interface to manage http operations of Pokemon entities
type Pokemon struct {
	ParseURL func(r *http.Request, i interface{}) error
}

// GetList get list of pokemons
func (p Pokemon) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetList")
}

// GetOne get specific pokemon by id
func (p Pokemon) GetOne(w http.ResponseWriter, r *http.Request) {
	var param ParamsURL
	if err := p.ParseURL(r, &param); err != nil {
		rest.Fail(w, "", "Error parsing url", err)
	}
	pokemon := model.Pokemon{Id: param.Id, Name: "Pikachu"}
	p.Render(w, pokemon)
}

// Render renders to specific interface of Render
func (p Pokemon) Render(w http.ResponseWriter, i interface{}) {
	if err := (render.JsonRender{Data: i}.Render(w)); err != nil {
		panic(err)
	}
}
