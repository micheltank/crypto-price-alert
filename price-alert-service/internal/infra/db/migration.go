package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/infra/config"
	"github.com/pkg/errors"
)

func Migrate(dbConfig config.DbConfig) error {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return errors.Wrap(err, "failed to open postgres connection")
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to instantiate postgres driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dbConfig.Name, driver)
	if err != nil {
		return errors.Wrap(err, "failed to create migrate instance")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to apply migrate up")
	}
	return nil
}
