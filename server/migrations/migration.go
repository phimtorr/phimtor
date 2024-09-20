package migrations

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrationFS embed.FS

func Run(db *sql.DB) error {
	d, err := iofs.New(migrationFS, ".")
	if err != nil {
		return fmt.Errorf("create iofs source instance: %w", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("create mysql driver instance: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", d, "mysql", driver)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
