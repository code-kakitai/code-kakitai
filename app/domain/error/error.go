package error

import "errors"

type Error struct {
	description string
}

func (e *Error) Error() string {
	return e.description
}

func NewError(s string) *Error {
	return &Error{
		description: s,
	}
}

func IsValidationError(err error) error {
	var e *Error
	if errors.As(err, &e) {
		return e
	}
	return err
}

var NotFoundErr = errors.New("not found")
