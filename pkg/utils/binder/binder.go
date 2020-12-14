package binder

import "io"

// Binder is the interface to wraps different binders
type Binder interface {
	Binder(r io.Reader) error
}
