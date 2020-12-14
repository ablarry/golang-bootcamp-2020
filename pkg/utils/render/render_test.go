package render

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJSONRender
func TestJSONRender(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"id":   "1",
		"name": "pikachu",
	}
	err := (JSONRender{data}).Render(w)
	assert.NoError(t, err)
	assert.Equal(t, "{\"id\":\"1\",\"name\":\"pikachu\"}", w.Body.String())
}
