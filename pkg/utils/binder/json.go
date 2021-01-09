package binder

import (
	"encoding/json"
	"io"

	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/validator"
)

// JSON implements Binder interface
type JSON struct {
}

// Bind bind json param o json to dst
func (b JSON) Bind(r io.Reader, dst interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(dst); err != nil {
		return errors.NewError(errors.Payload, "Error: Invalid Payload", err)
	}
	return validator.Validate(dst)
}
