package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/ablarry/golang-bootcamp-2020/pkg/pokemons/api"
	
	"github.com/gorilla/mux"
)

func TestApi(t *testing.T) {
	r := mux.NewRouter()
	r = api.Initialize(r)
	ts := httptest.NewServer(r)
	defer ts.Close()
	type request struct {
		name     string
		endpoint string
		method   string
	}
	requests := []request{
		{"Invoke /pokemons", ts.URL + "/pokemons", http.MethodGet},
		{"Invoke /pokemons/1", ts.URL + "/pokemons/1", http.MethodGet},
	}
	for _, v := range requests {
		t.Run(v.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(v.method, v.endpoint, nil)
			if err != nil {
				log.Fatal(err)
			}
			_, err = client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
		})

	}
}
