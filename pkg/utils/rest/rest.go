package rest

import (
	"fmt"
	"net/http"

	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/binder"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/render"
)

// JSON writes JSON ContentType.
func JSON(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if i != nil {
		if err := (render.JSONRender{Data: i}.Render(w)); err != nil {
			panic(err)
		}
	}
}

// BindJSON binds http.Request.body JSON ContentType
func BindJSON(r *http.Request, i interface{}) error {
	b := &binder.JSON{}
	return b.Bind(r.Body, i)
}

// statusMap map to http error
func statusMap() map[string]int {
	m := map[string]int{
		errors.Payload:   http.StatusBadRequest,
		errors.NotFound:  http.StatusNotFound,
		errors.Duplicate: http.StatusUnprocessableEntity,
	}
	return m
}

// Fail convenience for REST error responses.
func Fail(w http.ResponseWriter, code, message string, err error) {
	res := &errors.Error{
		Code:    code,
		Message: message,
		Detail:  err.Error(),
	}
	Err, ok := err.(*errors.Error)
	if ok {
		res = Err
	}
	sc, ok := statusMap()[res.Code]
	if !ok {
		sc = http.StatusInternalServerError
	}
	w.WriteHeader(sc)
	if err = render.Write(w, res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error rendering response")
	}
}

// Handler convenience for REST error responses.
func Handler(w http.ResponseWriter, err error) {
	var res *errors.Error
	Err, ok := err.(*errors.Error)
	if ok {
		res = Err
	} else {
		res = errors.NewError("500", "Error server", err)
	}
	sc, ok := statusMap()[res.Code]
	if !ok {
		sc = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sc)
	if err = render.Write(w, res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error rendering response")
	}
}
