package binder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJSONBind binding json
func TestJSONBinding(t *testing.T) {
	var s struct {
		Name string `json:"name"`
	}
	err := JSON{}.Bind(bytes.NewReader([]byte(`{"name": "pikachu"}`)), &s)
	assert.NoError(t, err, "Error not expected bindin")
	assert.Equal(t, "pikachu", s.Name)
}
