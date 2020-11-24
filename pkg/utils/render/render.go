package render

import "net/http"

type Render interface {
	Render(w http.ResponseWriter) (err error)
}
