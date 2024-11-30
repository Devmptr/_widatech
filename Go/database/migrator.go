package database

import (
	"widatech_interview/golang/config"
	"widatech_interview/golang/core"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Execute(conf *config.DatabaseConfig) error {
	var err error

	db := core.NewDB(conf)
	if err := db.MakeConnection(); err != nil {
		return err
	}
	defer db.CloseConnection()

	driver, err := mysql.WithInstance(db.Connection, &mysql.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
