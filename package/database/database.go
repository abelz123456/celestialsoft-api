package database

import (
	"context"
	"fmt"
	"time"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Driver         DBDriver
	Address        string
	ConnectionInfo string

	Sql   *gorm.DB
	Mongo *mongo.Database
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

		db.Sql, err = gorm.Open(mysql.Open(db.Address), &gorm.Config{})
		if err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

		sqlDB, err := db.Sql.DB()
		if err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

		if err := sqlDB.PingContext(ctx); err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

	case PostgreSQL:
		db.Address = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.PostgresqlDBUser, cfg.PostgreqslDBPass, cfg.PostgresqlDBHost, cfg.PostgresqlDBPort, cfg.PostgresqlDBName)
		db.ConnectionInfo = fmt.Sprintf("postgresql(%s:%s/%s)", cfg.MysqlDBHost, cfg.MysqlDBPort, cfg.MysqlDBName)

		db.Sql, err = gorm.Open(postgres.Open(db.Address))
		if err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

		sql, err := db.Sql.DB()
		if err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

		if err := sql.PingContext(ctx); err != nil {
			logger.Error(err, "", nil)
			return nil, err
		}

	case Mongo:
		db.Address = fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoDBUser, cfg.MongoDBPass, cfg.MongoDBHost, cfg.MongoDBPort)
		if cfg.MongoDBPort == 0 {
			db.Address = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", cfg.MongoDBUser, cfg.MongoDBPass, cfg.MongoDBHost)
		}

		db.ConnectionInfo = fmt.Sprintf("mongodb(%s:%d)", cfg.MongoDBHost, cfg.MongoDBPort)
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		mongoconn := options.Client().ApplyURI(db.Address).
			SetServerAPIOptions(serverAPI)
		client, _ := mongo.Connect(ctx, mongoconn)
		if err = client.Ping(ctx, nil); err != nil {
			logger.Error(err, "failed while connecting database", nil)
			return nil, err
		}

		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			logger.Error(err, "", nil)
		}
		fmt.Println(databases)

		db.Mongo = client.Database(cfg.MongoDBName)

	default:
		err = fmt.Errorf("\"%s\" is invalid value. $DB_USED only use mysql/postgresql/mongodb", cfg.DBUsed)
		logger.Error(err, "", nil)
		return nil, err
	}

	ctx.Done()
	logger.Info("Database connected!", map[string]interface{}{"address": db.ConnectionInfo})
	return &db, nil
}
