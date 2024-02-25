package migrations

import (
	"database/sql"
	"embed"

	"github.com/friendsofgo/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrationFS embed.FS

func Run(db *sql.DB) error {
	d, err := iofs.New(migrationFS, ".")
	if err != nil {
		return errors.Wrap(err, "create iofs source instance")
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return errors.Wrap(err, "create mysql driver instance")
	}
	m, err := migrate.NewWithInstance("iofs", d, "mysql", driver)
	if err != nil {
		return errors.Wrap(err, "create migrate instance")
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "run migrations")
	}
	return nil
}
