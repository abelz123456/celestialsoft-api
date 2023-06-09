package config

type Config struct {
	BasePath        string `mapstructure:"-"`
	AppName         string `mapstructure:"APP_NAME"`
	AppEnv          string `mapstructure:"APP_ENV"`
	DevelopmentPort string `mapstructure:"DEVELOPMENT_PORT"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	JwtExpiredTime  int    `mapstructure:"JWT_EXPIRED_TIME"`
	AppHost         string `mapstructure:"APP_HOST"`
	AppScheme       string `mapstructure:"APP_SCHEME"`
	StaticFilePath  string `mapstructure:"STATIC_FILE_PATH"`

	TrustedProxies []string `mapstructure:"TRUSTED_PROXIES"`

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
	PostgresqlDBName string `mapstructure:"POSTGRESQL_DB_NAME"`

	MongoDBHost string `mapstructure:"MONGO_DB_HOST"`
	MongoDBPort int    `mapstructure:"MONGO_DB_PORT"`
	MongoDBUser string `mapstructure:"MONGO_DB_USER"`
	MongoDBPass string `mapstructure:"MONGO_DB_PASS"`
	MongoDBName string `mapstructure:"MONGO_DB_NAME"`

	RajaongkirApiKey string `mapstructure:"RAJAONGKIR_API_KEY"`
	RajaongkirApiUrl string `mapstructure:"RAJAONGKIR_API_URL"`

	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpUser     string `mapstructure:"SMTP_USER"`
	SmtpPass     string `mapstructure:"SMTP_PASS"`
	MailFromAddr string `mapstructure:"MAIL_FROM_ADDR"`
	MailFromName string `mapstructure:"MAIL_FROM_NAME"`
}
