package database

import (
	"testing"

	"github.com/abelz123456/celestial-api/package/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDatabaseMySQLSuccess(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = MySQL.String()

	db, err := NewDatabase(cfg)

	require.NoError(t, err)
	assert.Equal(t, MySQL, db.Driver)
	assert.Contains(t, db.Address, cfg.MysqlDBHost)
	assert.Contains(t, db.ConnectionInfo, cfg.MysqlDBHost)
	assert.NotNil(t, db.Sql)
}

func TestNewDatabaseMySQLFailure(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = MySQL.String()
	cfg.MysqlDBHost = "local.host"

	db, err := NewDatabase(cfg)

	require.Error(t, err)
	assert.Nil(t, db)
}

func TestNewDatabasePostgreSQLSuccess(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = PostgreSQL.String()

	db, err := NewDatabase(cfg)

	require.NoError(t, err)
	assert.Equal(t, PostgreSQL, db.Driver)
	assert.Contains(t, db.Address, cfg.MysqlDBHost)
	assert.Contains(t, db.ConnectionInfo, cfg.MysqlDBHost)
	assert.NotNil(t, db.Sql)
}

func TestNewDatabasePostgreSQLFailure(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = PostgreSQL.String()
	cfg.PostgresqlDBHost = "local.host"

	db, err := NewDatabase(cfg)

	require.Error(t, err)
	assert.Nil(t, db)
}

func TestNewDatabaseMongoDBSuccess(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = Mongo.String()

	db, err := NewDatabase(cfg)

	require.NoError(t, err)
	assert.Equal(t, Mongo, db.Driver)
	assert.Contains(t, db.Address, cfg.MysqlDBHost)
	assert.Contains(t, db.ConnectionInfo, cfg.MysqlDBHost)
	assert.NotNil(t, db.Mongo)
}

func TestNewDatabaseMongoDBFailure(t *testing.T) {
	cfg := config.Init(".")
	cfg.DBUsed = Mongo.String()
	cfg.MongoDBHost = "local.host"

	db, err := NewDatabase(cfg)

	require.Error(t, err)
	assert.Nil(t, db)
}

func TestNewDatabaseInvalidDriver(t *testing.T) {
	cfg := config.Config{DBUsed: "mssql"}

	db, err := NewDatabase(cfg)

	require.Error(t, err)
	assert.Nil(t, db)
}
