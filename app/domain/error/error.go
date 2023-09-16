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

var NotFoundErr = errors.New("not found")
