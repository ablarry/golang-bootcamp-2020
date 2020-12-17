package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidate implement validations for structs and individual fields based on tags.
func TestValidate(t *testing.T) {
	table := []struct {
		name string
		form interface{}
		ok   bool
	}{
		{
			"1._ OK Validate field required",
			struct {
				Email string `validate:"required"`
			}{"mail@mail.com"},
			true,
		},

		{
			"2._ Error Validate field required",
			struct {
				Email string `validate:"required"`
			}{},
			false,
		},
		{
			"3._ Ok Validate field not required",
			struct {
				Email string `form:":mail"`
			}{},
			true,
		},
		{
			"4._ Ok Validate integer not required",
			struct {
				Size int64 `form:":size"`
			}{6},
			true,
		},

		{
			"5._ Ok Validate integer not required",
			struct {
				Size int64 `form:":size"`
			}{},
			true,
		},
	}

	for _, tt := range table {
		err := Validate(tt.form)
		if tt.ok {
			assert.NoError(t, err, "Error validation not expected")
		} else {
			assert.Error(t, err, "Error expected")
		}

	}
}
