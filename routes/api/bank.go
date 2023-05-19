package api

import (
	"github.com/abelz123456/celestial-api/api/bank/controllers"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/package/middleware"
	"github.com/gin-gonic/gin"
)

func NewBankApi(route *gin.RouterGroup, mgr manager.Manager) {
	ctrlr := controllers.NewController(mgr)

	route.GET("/", middleware.AuthenticatedMiddleware(mgr), ctrlr.GetList)
	route.POST("/", middleware.AuthenticatedMiddleware(mgr), ctrlr.CreateNew)
	route.GET("/:oid", middleware.AuthenticatedMiddleware(mgr), ctrlr.GetOne)
	route.PATCH("/:oid", middleware.AuthenticatedMiddleware(mgr), ctrlr.UpdateOne)
	route.DELETE("/:oid", middleware.AuthenticatedMiddleware(mgr), ctrlr.Delete)
}
