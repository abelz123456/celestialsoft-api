package repositories

import (
	"github.com/abelz123456/celestial-api/api/auth/domain"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/manager"
)

func NewRepository(mgr manager.Manager) domain.Repository {
	switch mgr.Config.DBUsed {
	case database.MySQL.String():
		return &mysql{
			Sql: mgr.Database.Sql,
			Log: mgr.Logger,
		}
	case database.PostgreSQL.String():
		return &postgresql{
			Sql: mgr.Database.Sql,
			Log: mgr.Logger,
		}
	default:
		return &mongodb{
			Mongo: mgr.Database.Mongo,
			Log:   mgr.Logger,
		}
	}
}
