package postgres

import (
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (r DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	uri, err := getPostgresConnectionURI()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open(postgresDriverName, uri)

	if err != nil {
		log.Print(postgresConnectionErrorMsg + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func MigrateUp() error {
	connectionURI, err := getPostgresConnectionURI()
	if err != nil {
		return err
	}

	db, err := sql.Open(postgresDriverName, connectionURI)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./config/sql/migrations/", databaseName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return db.Close()
}

func (r DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Error().Err(err)
	}
}
