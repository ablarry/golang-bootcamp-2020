package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	// val creation of validator
	val = validator.New()
)

// Validate validate v structure according tags defined in fields
func Validate(v interface{}) error {

	// returns nil or ValidationErrors ( []FieldError )
	err := val.Struct(v)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}
	return nil
}
