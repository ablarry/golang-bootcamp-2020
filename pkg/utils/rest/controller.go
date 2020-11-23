package rest

import "net/http"

// Controller is a general interface to manage http operations
type Controller interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
}
