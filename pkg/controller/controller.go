package controller

import (
	"log"
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/service"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/rest"
)

// Pokemon is a specific interface to manage http operations of Pokemon entities
type Pokemon struct {
	S service.Service
}

// GetList get list of pokemons
func (c Pokemon) GetList(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GetList")
	params := ParamsList{Size: 10, Paginator: 0}
	if err := rest.DecodeURL(r, &params); err != nil {
		rest.Handler(w, err)
		return
	}
	pk, err := c.S.GetList(params.Paginator, params.Size)
	if err != nil {
		rest.Handler(w, err)
		return
	}
	rest.JSON(w, pk, http.StatusOK)

}

// GetOne get specific pokemon by id
func (c Pokemon) GetOne(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GetOne")
	var param ParamsOne
	if err := rest.DecodeURL(r, &param); err != nil {
		rest.Handler(w, err)
		return
	}
	pk, err := c.S.GetOne(param.ID)
	if err != nil {
		rest.Handler(w, err)
		return
	}
	rest.JSON(w, pk, http.StatusOK)
}

// Create creates a register
func (c Pokemon) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Create")
	pk := &model.Pokemon{}
	err := rest.BindJSON(r, pk)
	if err != nil {
		rest.Handler(w, err)
		return
	}
	_, err = c.S.Create(pk)
	if err != nil {
		rest.Handler(w, err)
		return
	}
	rest.JSON(w, pk, http.StatusCreated)
}

// Delete deletes a register
func (c Pokemon) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Create")
	var param ParamsOne
	if err := rest.DecodeURL(r, &param); err != nil {
		rest.Handler(w, err)
		return
	}
	rows, err := c.S.Delete(param.ID)
	if err != nil {
		rest.Handler(w, err)
		return
	}
	if rows == 0 {
		rest.JSON(w, nil, http.StatusNoContent)
		return
	}
	rest.JSON(w, nil, http.StatusOK)
}
