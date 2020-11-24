package rest

import (
	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// DecodeURL decodes URL param into i
func DecodeURL(r *http.Request, i interface{}) error{
	decoder := form.NewDecoder()
	v := r.URL.Query()
	log.Printf("DecodeURL : url vars : %v : %v : %v", v, mux.Vars(r), r.RequestURI)
	// Add path params as values.
	for k, x := range mux.Vars(r) {
		v.Set(":"+k, x)
	}
	if err := decoder.Decode(i,v); err != nil {
		return NewError(Payload, "Error decode", err)
	}
	return nil
}
