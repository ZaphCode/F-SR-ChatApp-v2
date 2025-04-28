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
	messages := map[string]string{
		"required": "This field is required",
		"lte":      "Should be less than " + fe.Param(),
		"gte":      "Should be greater than " + fe.Param(),
		"max":      "Should be less than " + fe.Param() + " characters",
		"min":      "Should be greater than " + fe.Param() + " characters",
		"email":    "Invalid email",
		"oneof":    "Yo can only chose between: [" + fe.Param() + "]",
		"url":      "Invalid url",
	}

	if msg, ok := messages[fe.Tag()]; ok {
		return msg
	}

	return fmt.Sprintf("Unknown error (%s)", fe.Tag())
}
