package api

import (
	"github.com/abelz123456/celestial-api/api/mail/controllers"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/package/middleware"
	"github.com/gin-gonic/gin"
)

func NewMailApi(route *gin.RouterGroup, mgr manager.Manager) {
	ctrl := controllers.NewController(mgr)

	route.POST("/", middleware.AuthenticatedMiddleware(mgr), ctrl.Send)
	route.GET("/", middleware.AuthenticatedMiddleware(mgr), ctrl.EmailSentCollection)
	route.GET("/:uid", middleware.AuthenticatedMiddleware(mgr), ctrl.InfoByUID)
}
