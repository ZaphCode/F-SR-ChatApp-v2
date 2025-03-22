package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	return fmt.Sprintf("Validation errors: %v", []ValidationError(e))
}

func validate(dto any) error {
	validate := validator.New()

	if err := validate.Struct(dto); err != nil {
		var errors ValidationErrors
		for _, fe := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field: fe.Field(),
				Error: getErrorMsg(fe),
			})
		}
		return errors
	}

	return nil
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "max":
		return "Should be less than " + fe.Param() + " characters"
	case "min":
		return "Should be greater than " + fe.Param() + " characters"
	case "email":
		return "Invalid email"
	case "oneof":
		return fmt.Sprintf("Yo can only chose between: [%s]", fe.Param())
	case "url":
		return "Invalid url"
	default:
		return fmt.Sprintf("Unknown error (%s)", fe.Tag())
	}
}
