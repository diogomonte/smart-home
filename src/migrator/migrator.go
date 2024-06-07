package migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
}

func (m *Migrator) RunMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %v", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://migrations",
		"mysql",
		driver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	defer func() {
		migrator.Close()
	}()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations %v", err)
	}

	return nil
}
