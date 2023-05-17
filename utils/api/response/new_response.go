package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse string
type responseProperties struct {
	ResultCode string      `json:"resultCode"`
	HttpStatus int         `json:"httpStatus"`
	Message    string      `json:"message,omitempty"`
	Total      int         `json:"total,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func getProperties(r ApiResponse) (responseProperties, bool) {
	responseMap := make(map[ApiResponse]responseProperties)
	noProperty := responseProperties{
		ResultCode: "000",
		HttpStatus: http.StatusFailedDependency,
		Message:    "Response %s is not in responseMap",
	}

	for key, val := range successResponseMap {
		responseMap[key] = val
	}

	for key, val := range failedResponseMap {
		responseMap[key] = val
	}

	prop, exixts := responseMap[r]
	if exixts {
		return prop, true
	}

	return noProperty, false
}

func SendJson(ctx *gin.Context, apiResponse ApiResponse, message string, data interface{}) {
	property, _ := getProperties(apiResponse)
	property.Data = nil

	if message != "" {
		property.Message = message
	}

	if _, ok := data.(error); ok {
		data = data.(error).Error()
	}

	if data != nil {
		property.Data = data
	}

	ctx.JSON(property.HttpStatus, property)
	ctx.Abort()
}
