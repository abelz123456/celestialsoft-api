package server

import (
	"net/http"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/server/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	logger     log.Log
	Engine     *gin.Engine
	HttpServer *http.Server
}

func Init(cfg config.Config) (*Server, error) {
	var (
		port   = cfg.DevelopmentPort
		engine = gin.Default()
	)

	engine.MaxMultipartMemory = 8 << 20 // maximum file 8 MiB

	// Load Middleware
	engine.Use(middleware.CORSMiddleware)
	engine.Use(gin.Logger())
	engine.Use(middleware.RecoveryMiddleware())
	engine.SetTrustedProxies(cfg.TrustedProxies)

	return &Server{
		logger: log.NewLog(),
		Engine: engine,
		HttpServer: &http.Server{
			Addr:    port,
			Handler: engine,
		},
	}, nil
}
