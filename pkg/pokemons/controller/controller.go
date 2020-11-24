package controller

import (
	"fmt"
	"github.com/ablarry/golang-bootcamp-2020/pkg/pokemons/model"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/render"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/rest"
	"net/http"
)

type Controller interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
}

type Pokemon struct {
	ParseURL func (r *http.Request, i interface{}) error
}



func (p Pokemon) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetList")
}

func (p Pokemon) GetOne(w http.ResponseWriter, r *http.Request) {
	var param ParamsURL
	if err := p.ParseURL(r, &param); err != nil {
		rest.Fail(w,"","",err)
	}
	pokemon := model.Pokemon{Id: param.Id ,Name: "Pikachu"}
	render.JsonRender{Data: pokemon}.Render(w)
}
