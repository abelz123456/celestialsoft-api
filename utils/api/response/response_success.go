package response

import "net/http"

const (
	Ok      ApiResponse = "20000"
	Created ApiResponse = "20101"
)

var successResponseMap = map[ApiResponse]ResponseProperties{
	Ok: {
		ResultCode: string(Ok),
		HttpStatus: http.StatusOK,
		Message:    "Ok",
	},
	Created: {
		ResultCode: string(Created),
		HttpStatus: http.StatusCreated,
		Message:    "New Data Created",
	},
}
