package postgres

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	postgresDriverName         = "postgres"
	postgresConnectionErrorMsg = "Error while accessing database: "
	noDataFound                = "sql: no rows in result set"
	passwordsDoNotMatch        = "invalid password for the user"
	databaseName               = "colabora_db"
)
