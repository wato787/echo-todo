package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using validator tags
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, formatValidationError(err))
		}
		return errors.New(strings.Join(validationErrors, ", "))
	}
	return nil
}

func formatValidationError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())
	switch err.Tag() {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + err.Param() + " characters long"
	case "max":
		return field + " must be at most " + err.Param() + " characters long"
	case "email":
		return field + " must be a valid email address"
	default:
		return field + " is invalid"
	}
}