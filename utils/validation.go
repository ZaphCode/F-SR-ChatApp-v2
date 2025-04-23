package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type validationErrors map[string]string

func (e validationErrors) Error() string {
	return fmt.Sprintf("validation errors: %v", validationErrors(e))
}

func Validate(dto any) error {
	validate := validator.New()
	errors := validationErrors{}

	if err := validate.Struct(dto); err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			errors[fe.Field()] = getErrorMsg(fe)
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
