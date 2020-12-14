package render

import (
	"encoding/json"
	"io"
)

// JSONRender struct is a general wrapper response body
type JSONRender struct {
	Data interface{}
}

// Render implements from general interface Render
func (r JSONRender) Render(w io.Writer) (err error) {
	if err = Write(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// Write converts the obj argument to json
func Write(w io.Writer, obj interface{}) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
