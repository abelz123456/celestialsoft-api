package helpers

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidationErrorToMap(t *testing.T) {
	// Create a new validator instance.
	v := validator.New()

	// Create a struct with some validation errors.
	type User struct {
		Name string `validate:"required"`
	}
	u := User{Name: ""}

	// Validate the struct.
	err := v.Struct(u)

	// Assert that the error is not nil.
	assert.NotNil(t, err)

	// Convert the error to a map.
	errorMap := ValidationErrorToMap(err)

	assert.NotNil(t, errorMap)

	// Assert that the map contains the expected error messages.
	assert.Equal(t, len(*errorMap), 1)
	mapError := *errorMap

	assert.Contains(t, mapError["Name"], "Error:Field validation for 'Name' failed on the 'required' tag")
}

func TestValidationErrorToMapWithInvalidError(t *testing.T) {
	err := errors.New("invalid error")

	// Convert the error to a map.
	errorMap := ValidationErrorToMap(err)

	assert.Nil(t, errorMap)
}
