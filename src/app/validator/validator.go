package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(i interface{}) error {
	err := validate.Struct(i)
	return err
}
