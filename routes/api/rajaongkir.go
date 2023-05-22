package api

import (
	"github.com/abelz123456/celestial-api/api/rajaongkir/controllers"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/package/middleware"
	"github.com/gin-gonic/gin"
)

func NewRajaongkirApi(route *gin.RouterGroup, mgr manager.Manager) {
	ctrl := controllers.NewController(mgr)

	route.GET("/", middleware.AuthenticatedMiddleware(mgr), ctrl.GetHistories)
	route.GET("/province", middleware.AuthenticatedMiddleware(mgr), ctrl.GetProvince)
	route.GET("/province/:id/city", middleware.AuthenticatedMiddleware(mgr), ctrl.GetCity)
	route.POST("/cost", middleware.AuthenticatedMiddleware(mgr), ctrl.CostInfo)
}
