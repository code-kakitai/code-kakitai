package errors

import (
	"github.com/cockroachdb/errors"
)

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

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func WithStack(err error) error {
	return errors.WithStack(err)
}
