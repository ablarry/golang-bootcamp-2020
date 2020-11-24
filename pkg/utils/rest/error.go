package rest

import (
	"fmt"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/render"
	"net/http"
)
const (

	// Payload error.
	Payload = "validation.payload"

)
func statusMap() map[string]int {
	m := map[string]int{
		Payload:      http.StatusBadRequest,

	}
	return m
}

// Error REST type.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// NewError helper for new error.
func NewError(code, message string, err error) *Error {
	e := &Error{
		Code:    code,
		Message: message,
		Detail:  fmt.Sprintf("%s", err),
	}
	return e
}

// Error implements error
func (e *Error) Error() string {
	return e.Detail
}



// Fail convenience for REST error responses.
func Fail(w  http.ResponseWriter, code, message string, err error) {
	res := &Error{
		Code:    code,
		Message: message,
		Detail:  err.Error(),
	}
	Err, ok := err.(*Error)
	if ok {
		res = Err
	}
	sc, ok := statusMap()[res.Code]
	if !ok {
		sc = http.StatusInternalServerError
	}
	w.WriteHeader(sc)
	render.Write(w,Err)
}