package api

import (
	"github.com/abelz123456/celestial-api/api/file/controllers"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/package/middleware"
	"github.com/gin-gonic/gin"
)

func NewFileApi(route *gin.RouterGroup, mgr manager.Manager) {
	ctrlr := controllers.NewController(mgr)

	route.POST("/", middleware.AuthenticatedMiddleware(mgr), ctrlr.Upload)
	route.GET("/", middleware.AuthenticatedMiddleware(mgr), ctrlr.GetCollection)
	route.GET("/:uid", middleware.AuthenticatedMiddleware(mgr), ctrlr.GetInfo)
	route.PUT("/:uid", middleware.AuthenticatedMiddleware(mgr), ctrlr.Replace)
	route.DELETE("/:uid", middleware.AuthenticatedMiddleware(mgr), ctrlr.Unlink)
}
