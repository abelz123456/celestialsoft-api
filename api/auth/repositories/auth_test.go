package repositories

import (
	"testing"

	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/test/mockdata"
)

func TestNewMySQLRepository(t *testing.T) {
	mysqlMgr := mockdata.NewFakeManager(t, database.MySQL)
	NewRepository(mysqlMgr)
}

func TestNewPostgreSQLRepository(t *testing.T) {
	postgresqlMgr := mockdata.NewFakeManager(t, database.PostgreSQL)
	NewRepository(postgresqlMgr)
}

func TestNewMongoDBRepository(t *testing.T) {
	mongodbMgr := mockdata.NewFakeManager(t, database.Mongo)
	NewRepository(mongodbMgr)
}
