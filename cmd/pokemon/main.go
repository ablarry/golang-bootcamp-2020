package main

import (
	"log"
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/api"
	"github.com/ablarry/golang-bootcamp-2020/pkg/config"
	"github.com/ablarry/golang-bootcamp-2020/pkg/infra"

	"github.com/gorilla/mux"
)

func main() {
	conf, err := config.Load("app")
	if err != nil {
		log.Fatal(err)
	}
	db, err := infra.InitDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r = api.Initialize(r, db)

	http.Handle("/", r)

	// Swagger ui dist
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/"))))

	log.Fatal(http.ListenAndServe(conf.Server.Port, nil))
}
