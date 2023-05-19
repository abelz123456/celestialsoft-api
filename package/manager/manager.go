package manager

import (
	"github.com/abelz123456/celestial-api/docs"
	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/server"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Manager struct {
	Logger   log.Log
	Config   config.Config
	Server   server.Server
	Database database.Database
}

func Init(path string) (mgr *Manager, err error) {
	var (
		cfg = config.Init(path)
	)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		return mgr, err
	}

	server, err := server.Init(cfg)
	if err != nil {
		return mgr, err
	}

	initSwagger(cfg, *server)

	return &Manager{
		Logger:   log.NewLog(),
		Config:   cfg,
		Database: *db,
		Server:   *server,
	}, nil
}

func initSwagger(cfg config.Config, server server.Server) {
	engine := server.Engine

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Host = cfg.AppHost
	docs.SwaggerInfo.Schemes = []string{cfg.AppScheme}
}
