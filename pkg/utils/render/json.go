package render

import (
	"encoding/json"
	"net/http"
)

// JsonRender struct is a general wrapper response body
type JSONRender struct {
	Data interface{} `json:"data"`
}

// Render implements from general interface Render
func (r JsonRender) Render(w http.ResponseWriter) (err error) {
	if err = Write(w, r); err != nil {
		panic(err)
	}
	return
}

// Write converts the obj argument to json
func Write(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
