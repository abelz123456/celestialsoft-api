package manager

import (
	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/server"
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

	return &Manager{
		Logger:   log.NewLog(),
		Config:   cfg,
		Database: *db,
		Server:   *server,
	}, nil
}
