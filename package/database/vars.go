package database

type DBDriver string

func (v DBDriver) String() string {
	return string(v)
}

var (
	MySQL      DBDriver = "mysql"
	PostgreSQL DBDriver = "postgres"
	Mongo      DBDriver = "mongodb"
)
