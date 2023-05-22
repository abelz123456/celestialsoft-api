package response

import "net/http"

const (
	ErrBadRequest              ApiResponse = "40000"
	ErrInternalServer          ApiResponse = "40001"
	ErrFailedRegister          ApiResponse = "40002"
	ErrFailedGetBankCollection ApiResponse = "40003"
	ErrFailedCreateBank        ApiResponse = "40004"
	ErrFailedRemoveBank        ApiResponse = "40005"
	ErrFailedGetBank           ApiResponse = "40006"
	ErrFailedUpdateBank        ApiResponse = "40007"
	ErrFailedUploadFile        ApiResponse = "40008"
	ErrFailedGetCost           ApiResponse = "40009"
	ErrFailedSendEmail         ApiResponse = "40010"
	ErrFailedLogin             ApiResponse = "40100"
	ErrUnauthorized            ApiResponse = "401001"
	ErrForm1Forbidden          ApiResponse = "40300"
)

var failedResponseMap = map[ApiResponse]ResponseProperties{
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
	ErrFailedGetBankCollection: {
		ResultCode: string(ErrFailedGetBankCollection),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle get bank collection",
	},
	ErrFailedCreateBank: {
		ResultCode: string(ErrFailedCreateBank),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle create bank",
	},
	ErrFailedRemoveBank: {
		ResultCode: string(ErrFailedRemoveBank),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle remove bank",
	},
	ErrFailedGetBank: {
		ResultCode: string(ErrFailedGetBank),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle get bank",
	},
	ErrFailedUpdateBank: {
		ResultCode: string(ErrFailedUpdateBank),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle update bank",
	},
	ErrFailedUploadFile: {
		ResultCode: string(ErrFailedUploadFile),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle upload file",
	},
	ErrFailedGetCost: {
		ResultCode: string(ErrFailedGetCost),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle get cost info",
	},
	ErrFailedSendEmail: {
		ResultCode: string(ErrFailedSendEmail),
		HttpStatus: http.StatusBadRequest,
		Message:    "Error handle send email",
	},
	ErrFailedLogin: {
		ResultCode: string(ErrFailedLogin),
		HttpStatus: http.StatusUnauthorized,
		Message:    "Error handle login",
	},
	ErrUnauthorized: {
		ResultCode: string(ErrUnauthorized),
		HttpStatus: http.StatusUnauthorized,
		Message:    "Invalid Auth Token",
	},
	ErrForm1Forbidden: {
		ResultCode: string(ErrForm1Forbidden),
		HttpStatus: http.StatusForbidden,
		Message:    "missing_or_invalid_parameter",
	},
}
