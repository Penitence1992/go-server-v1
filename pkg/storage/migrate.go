package storage

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const MigrationsQueryName = "x-migrations-table"

func DoMigrate(sourceUrl, dnsUrl string) error {
	m, err := migrate.New(sourceUrl, dnsUrl)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			_ = m.Down()
			return err
		}
		return nil
	}
	return nil
}
