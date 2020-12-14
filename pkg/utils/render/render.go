package render

import "net/http"

// Render is the interface to wraps different renders
type Render interface {
	Render(w http.ResponseWriter) (err error)
}
