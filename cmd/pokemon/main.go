package main

import (
	"log"
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/pokemons/api"
	
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r = api.Initialize(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
