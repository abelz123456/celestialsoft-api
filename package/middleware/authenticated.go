package middleware

import (
	"strings"

	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/api/response"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/gin-gonic/gin"
)

func AuthenticatedMiddleware(mgr manager.Manager) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			response.SendJson(ctx, response.ErrUnauthorized, "", nil)
			return
		}

		tokenSplited := strings.Split(authToken, "Bearer ")
		if len(tokenSplited) != 2 {
			response.SendJson(ctx, response.ErrUnauthorized, "", nil)
			return
		}

		jwt := helpers.NewJwtHelper(mgr.Config)
		sub := jwt.ParseToken(tokenSplited[1])
		if sub == "" {
			response.SendJson(ctx, response.ErrUnauthorized, "", nil)
			return
		}

		ctx.Set("oid", sub)
		ctx.Next()
	}
}
