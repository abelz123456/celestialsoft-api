package mockdata

import (
	"context"
	"testing"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/package/server"
	"github.com/stretchr/testify/assert"
)

func NewFakeManager(t *testing.T, dbDriver database.DBDriver) manager.Manager {
	var (
		db = database.Database{
			Driver: dbDriver,
		}
		logger = log.NewLog()
		cfg    = config.Config{
			DBUsed:          dbDriver.String(),
			DevelopmentPort: ":3030",
		}
	)

	switch dbDriver {
	case database.MySQL:
		_, db.Sql = NewMySQLMock(t)
	case database.PostgreSQL:
		_, db.Sql = NewPosgreSQLMock(t)
	default:
		mt := NewMongoDBMock(t, context.TODO())
		db.Mongo = mt.DB
	}

	s, err := server.Init(cfg)
	assert.NoError(t, err)

	return manager.Manager{
		Config:   cfg,
		Database: db,
		Logger:   logger,
		Server:   *s,
	}
}
