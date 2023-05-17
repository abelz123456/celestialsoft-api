package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidationErrorToMap(err error) *map[string]string {
	if _, ok := err.(validator.ValidationErrors); ok {
		validationErrors := err.(validator.ValidationErrors)
		errorMap := make(map[string]string)

		for _, e := range validationErrors {
			errorMap[e.Field()] = e.Error()
		}

		return &errorMap
	}

	return nil
}
