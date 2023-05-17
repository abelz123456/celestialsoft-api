package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Driver         DBDriver
	Address        string
	ConnectionInfo string

	Sql   *sql.DB
	Mongo *mongo.Client
}

func NewDatabase(cfg config.Config) (_ *Database, err error) {
	var (
		logger = log.NewLog()
		db     = Database{
			Driver: DBDriver(cfg.DBUsed),
		}
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch db.Driver {
	case MySQL:
		db.Address = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.MysqlDBUser, cfg.MysqlDBPass, cfg.MysqlDBHost, cfg.MysqlDBPort, cfg.MysqlDBName)
		db.ConnectionInfo = fmt.Sprintf("mysql(%s:%s/%s)", cfg.MysqlDBHost, cfg.MysqlDBPort, cfg.MysqlDBName)

		db.Sql, _ = sql.Open(db.Driver.String(), db.Address)
		if err := db.Sql.PingContext(ctx); err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

	case PostgreSQL:
		db.Address = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.PostgresqlDBUser, cfg.PostgreqslDBPass, cfg.PostgresqlDBHost, cfg.PostgresqlDBPort, cfg.PostgresqlDBName)
		db.ConnectionInfo = fmt.Sprintf("postgresql(%s:%s/%s)", cfg.MysqlDBHost, cfg.MysqlDBPort, cfg.MysqlDBName)

		db.Sql, _ = sql.Open(db.Driver.String(), db.Address)
		if err := db.Sql.PingContext(ctx); err != nil {
			logger.Error(err, db.Address, nil)
			return nil, err
		}

	case Mongo:
		db.Address = fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoDBUser, cfg.MongoDBPass, cfg.MongoDBHost, cfg.MongoDBPort)
		db.ConnectionInfo = fmt.Sprintf("mongodb(%s:%d)", cfg.MongoDBHost, cfg.MongoDBPort)
		mongoconn := options.Client().ApplyURI(db.Address)
		db.Mongo, _ = mongo.Connect(ctx, mongoconn)
		if err = db.Mongo.Ping(ctx, nil); err != nil {
			logger.Error(err, "failed while connecting database", nil)
			return nil, err
		}

	default:
		err = fmt.Errorf("\"%s\" is invalid value. $DB_USED only use mysql/postgresql/mongodb", cfg.DBUsed)
		logger.Error(err, "", nil)
		return nil, err
	}

	ctx.Done()
	logger.Info("Database connected!", map[string]interface{}{"address": db.ConnectionInfo})
	return &db, nil
}
