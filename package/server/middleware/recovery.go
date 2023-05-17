package middleware

import (
	errorHandler "github.com/abelz123456/celestial-api/utils/api/error_handler"
	"github.com/abelz123456/celestial-api/utils/api/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if !errorHandler.NotFoundErrors(c, recovered) && !errorHandler.ValidationErrors(c, recovered) {
			var message string = ""
			if _, ok := recovered.(*logrus.Entry); ok {
				message = recovered.(*logrus.Entry).Message
			} else if _, ok := recovered.(string); ok {
				message = recovered.(string)
			}

			response.SendJson(c, response.ErrInternalServer, "", message)
		}
	})
}
