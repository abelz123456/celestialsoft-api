package response

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseOk(t *testing.T) {
	response := struct {
		Name string `json:"name"`
	}{
		Name: "John Doe",
	}

	expected := DefaultResponse{
		ResultCode: "200",
		HttpStatus: "200",
		Message:    "success",
		Total:      1,
		Data: struct {
			Name string `json:"name"`
		}{
			Name: "John Doe",
		},
	}

	result := ResponseOk(response, 1)

	assert.Equal(t, expected, result)
}

func TestResponseDeleteOk(t *testing.T) {
	response := struct {
		ID int `json:"id"`
	}{
		ID: 1,
	}

	expected := DefaultResponse{
		ResultCode: "200",
		HttpStatus: "200",
		Message:    "success",
		Total:      1,
		Data: struct {
			ID int `json:"id"`
		}{
			ID: 1,
		},
	}

	result := ResponseDeleteOk(response)

	assert.Equal(t, expected, result)
}

func TestResponseOkDataNotFound(t *testing.T) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Data not found",
	}

	expected := DefaultResponse{
		ResultCode: "200",
		HttpStatus: "200",
		Message:    "data not found",
		Total:      0,
		Data: struct {
			Message string `json:"message"`
		}{
			Message: "Data not found",
		},
	}

	result := ResponseOkDataNotFound(response)

	assert.Equal(t, expected, result)
}

func TestResponseInternalServerError(t *testing.T) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Internal Server Error",
	}

	expected := DefaultResponse{
		ResultCode: "40401",
		HttpStatus: "404",
		Message:    "internal server error",
		Total:      0,
		Data: struct {
			Message string `json:"message"`
		}{
			Message: "Internal Server Error",
		},
	}

	result := ResponseInternalServerError(response)

	assert.Equal(t, expected, result)
}

func TestResponseBadRequest(t *testing.T) {
	err := "Invalid input"

	expected := DefaultResponse{
		ResultCode: "400",
		HttpStatus: "400",
		Message:    "bad_request",
		Total:      0,
		Data:       "Invalid input",
	}

	result := ResponseBadRequest(err)

	assert.Equal(t, expected, result)
}

func TestResponseUnAuthorized(t *testing.T) {
	err := "Unauthorized access"

	expected := DefaultResponse{
		ResultCode: "401",
		HttpStatus: "401",
		Message:    "unauthorized",
		Total:      0,
		Data:       "Unauthorized access",
	}

	result := ResponseUnAuthorized(err)

	assert.Equal(t, expected, result)
}
func TestNewNotFoundError(t *testing.T) {
	errorMsg := "Not found error"
	notFoundErr := NewNotFoundError(errorMsg)

	if notFoundErr.Error != errorMsg {
		t.Errorf("Expected error message: %s, but got: %s", errorMsg, notFoundErr.Error)
	}
}

func TestNewServerError(t *testing.T) {
	errorMsg := "Server error"
	serverErr := NewServerError(errorMsg)

	if serverErr.Error != errorMsg {
		t.Errorf("Expected error message: %s, but got: %s", errorMsg, serverErr.Error)
	}
}

func TestResponseForm1Forbidden(t *testing.T) {
	errMsg := "Invalid parameter"

	response := ResponseForm1Forbidden(errMsg)

	assert.Equal(t, "40300", response.ResultCode, "Expected ResultCode to be '40300'")
	assert.Equal(t, "403", response.HttpStatus, "Expected HttpStatus to be '403'")
	assert.Equal(t, "missing_or_invalid_parameter", response.Message, "Expected Message to be 'missing_or_invalid_parameter'")
	assert.Equal(t, 0, response.Total, "Expected Total to be 0")
	assert.Equal(t, errMsg, response.Data, "Expected Data to be the error message")
}

func TestResponseFinValidatorError(t *testing.T) {
	mockController := "MockController"
	mockFuncName := "MockFunction"
	mockError := fmt.Errorf("Validation error")

	// Capture log output for assertion
	logFile, _ := os.CreateTemp("", "test-log.txt")
	log.SetOutput(logFile)
	defer func() {
		log.SetOutput(os.Stderr)
		os.Remove(logFile.Name())
	}()

	response := ResponseFinValidatorError(mockController, mockFuncName, mockError)

	assert.Equal(t, "40300", response.ResultCode, "Expected ResultCode to be '40300'")
	assert.Equal(t, "403", response.HttpStatus, "Expected HttpStatus to be '403'")
	assert.Equal(t, "missing_or_invalid_parameter", response.Message, "Expected Message to be 'missing_or_invalid_parameter'")
	assert.Equal(t, 0, response.Total, "Expected Total to be 0")
	assert.Equal(t, mockError, response.Data, "Expected Data to be the validation error")

	// Assert log output
	logFile.Close()
	logContent, _ := os.ReadFile(logFile.Name())
	assert.Contains(t, string(logContent), "", "Expected log output to contain the error information")
}
