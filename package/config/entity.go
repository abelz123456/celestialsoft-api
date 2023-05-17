package config

type Config struct {
	BasePath        string `mapstructure:"-"`
	AppName         string `mapstructure:"APP_NAME"`
	AppEnv          string `mapstructure:"APP_ENV"`
	DevelopmentPort string `mapstructure:"DEVELOPMENT_PORT"`

	DBUsed string `mapstructure:"DB_USED"` // mysql | postgres | mongodb

	MysqlDBHost string `mapstructure:"MYSQL_DB_HOST"`
	MysqlDBPort string `mapstructure:"MYSQL_DB_PORT"`
	MysqlDBUser string `mapstructure:"MYSQL_DB_USER"`
	MysqlDBPass string `mapstructure:"MYSQL_DB_PASS"`
	MysqlDBName string `mapstructure:"MYSQL_DB_NAME"`

	PostgresqlDBHost string `mapstructure:"POSTGRESQL_DB_HOST"`
	PostgresqlDBPort string `mapstructure:"POSTGRESQL_DB_PORT"`
	PostgresqlDBUser string `mapstructure:"POSTGRESQL_DB_USER"`
	PostgreqslDBPass string `mapstructure:"POSTGRESQL_DB_PASS"`
	PostgresqlDBName string `'mapstructure:"POSTGRESQL_DB_NAME"`

	MongoDBHost string `mapstructure:"MONGO_DB_HOST"`
	MongoDBPort int    `mapstructure:"MONGO_DB_PORT"`
	MongoDBUser string `mapstructure:"MONGO_DB_USER"`
	MongoDBPass string `mapstructure:"MONGO_DB_PASS"`
}
