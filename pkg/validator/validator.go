package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func GetValidator() *validator.Validate {
	return validate
}

func NewValidator() *validator.Validate {
	return validate
}
