package database

import "database/sql"

type Properties struct {
	Driver         DBDriver
	Address        string
	ConnectionInfo string

	Sql *sql.DB
}
