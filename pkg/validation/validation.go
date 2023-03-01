package validation

import (
	"gopkg.in/go-playground/validator.v9"
)

func Validate(e interface{}) error {
	validate := validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
