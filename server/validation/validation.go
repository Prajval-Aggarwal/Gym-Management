package validation

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func CheckValidation(data interface{}) error {

	validationErr := Validate.Struct(data)
	if validationErr != nil {
		return validationErr
	}
	return nil
}
