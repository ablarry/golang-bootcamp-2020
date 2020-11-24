package render

import (
	"encoding/json"
	"net/http"
)

type JsonRender struct {
	Data interface{} `json:"data"`
}

func (r JsonRender) Render(w http.ResponseWriter) (err error) {
	if err = Write(w, r); err != nil {
		panic(err)
	}
	return
}

func Write(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
