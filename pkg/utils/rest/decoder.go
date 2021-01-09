package rest

import (
	"log"
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/validator"

	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
)

// DecodeURL decodes URL param into i
func DecodeURL(r *http.Request, i interface{}) error {
	decoder := form.NewDecoder()
	v := r.URL.Query()
	log.Printf("DecodeURL : url vars : %v : %v : %v", v, mux.Vars(r), r.RequestURI)
	// Add path params as values.
	for k, x := range mux.Vars(r) {
		v.Set(":"+k, x)
	}
	if err := decoder.Decode(i, v); err != nil {
		return errors.NewError(errors.Payload, "Error decode", err)
	}

	if err := validator.Validate(i); err != nil {

		return errors.NewError(errors.Payload, "Error validation", err)
	}
	return nil
}
