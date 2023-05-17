package api

import (
	"github.com/abelz123456/celestial-api/api/auth/controllers"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/gin-gonic/gin"
)

func NewAuthApi(route *gin.RouterGroup, mgr manager.Manager) {
	ctrlr := controllers.NewController(mgr)

	route.POST("/login", ctrlr.Login)
	route.POST("/register", ctrlr.Register)
}
