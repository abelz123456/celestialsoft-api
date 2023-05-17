package response

import "net/http"

const (
	ErrBadRequest     ApiResponse = "40000"
	ErrInternalServer ApiResponse = "40001"
	ErrFailedRegister ApiResponse = "40002"
	ErrForm1Forbidden ApiResponse = "40300"
)

var failedResponseMap = map[ApiResponse]responseProperties{
	ErrBadRequest: {
		ResultCode: string(ErrBadRequest),
		HttpStatus: http.StatusBadRequest,
		Message:    "General Error",
	},
	ErrInternalServer: {
		ResultCode: string(ErrInternalServer),
		HttpStatus: http.StatusBadRequest,
		Message:    "Internal Server Error",
	},
	ErrFailedRegister: {
		ResultCode: string(ErrFailedRegister),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle register",
	},
	ErrForm1Forbidden: {
		ResultCode: string(ErrForm1Forbidden),
		HttpStatus: http.StatusForbidden,
		Message:    "missing_or_invalid_parameter",
	},
}
