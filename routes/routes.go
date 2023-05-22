package routes

import (
	"net/http"

	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/routes/api"
	"github.com/gin-gonic/gin"
)

func LoadRoute(mgr manager.Manager) {
	var router = mgr.Server.Engine

	// Default Server PING
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
		ctx.Abort()
	})

	api.NewAuthApi(router.Group("/auth"), mgr)
	api.NewBankApi(router.Group("/bank"), mgr)
	api.NewFileApi(router.Group("/file"), mgr)
	api.NewMailApi(router.Group("/mail"), mgr)
	api.NewRajaongkirApi(router.Group("/rajaongkir"), mgr)
}
