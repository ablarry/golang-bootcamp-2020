package errors

	import "fmt"

const (

	// Payload error.
	Payload = "validation.payload"

	// NotFound error
	NotFound = "error.NotFound"

	// Duplicate error
	Duplicate = "error.Duplicate"
)

// ErrNotFound error NotFound
var ErrNotFound = NewError(NotFound, "Error: Register not found", nil)

// ErrDuplicate error Duplication
var ErrDuplicate = NewError(Duplicate, "Error: Register already exists", nil)

// ErrInvalidPayload error Payload
var ErrInvalidPayload = NewError(Payload, "Error: Invalid Payload", nil)

// Error REST type.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// NewError helper for new error.
func NewError(code, message string, err error) *Error {
	e := &Error{
		Code:    code,
		Message: message,
	}
	if err != nil {
		e.Detail = fmt.Sprintf("%s", err)
	}
	return e
}

// Error implements error
func (e *Error) Error() string {
	return e.Detail
}
