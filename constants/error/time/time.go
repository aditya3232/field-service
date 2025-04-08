package error

import "errors"

var (
	ErrTimeNotFound = errors.New("time not found")
)

var TimeErrors = []error{
	ErrTimeNotFound,
}
