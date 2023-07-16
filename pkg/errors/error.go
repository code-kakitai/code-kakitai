package errors

import (
	"github.com/cockroachdb/errors"
)

func NewError(s string) error {
	return errors.New(s)
}
