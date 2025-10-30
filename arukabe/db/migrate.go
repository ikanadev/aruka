package db

import (
	"arukabe/utils"

	// postgres connection and migration
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDB(config utils.Config) error {
	migrator, err := migrate.New(config.MigrationsSource, config.DBConn)
	if err != nil {
		return err
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
