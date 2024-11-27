package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

// Validate if a schema from a request is correct.
func ValidateSchema(payload interface{}) error {
	err := validate.Struct(payload)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return formatValidatorError(validationErrors)
	}
	return nil
}

func formatValidatorError(validationErr validator.ValidationErrors) error {
	var errorMessage string = ""
	for _, fieldError := range validationErr {
		errorMessage += fmt.Sprintf("Field '%s' failed in validations '%s'. value recived: '%v'",
			strings.ToLower(fieldError.Field()), fieldError.Tag(), fieldError.Value())
	}
	return fmt.Errorf("%s", errorMessage)
}
