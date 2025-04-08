package error

import "errors"

var (
	ErrFieldNotFound = errors.New("field not found")
)

var FieldErrors = []error{
	ErrFieldNotFound,
}
