package migrate

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_attempts = 10
	_timeout  = time.Second
)

func ApplyMigrations(migrationsPath string, dbURL string) {
	migrateLogger := log.New(os.Stdout, "[migrate] ", log.LstdFlags|log.Lshortfile)

	var m *migrate.Migrate
	var err error
	attempts := 0

	for attempts < _attempts {
		m, err = migrate.New(migrationsPath, dbURL)
		if err == nil {
			break
		}

		migrateLogger.Printf("Migrate: trying to connect, attempt %d", attempts+1)
		time.Sleep(_timeout)
		attempts++
	}

	if err != nil {
		migrateLogger.Fatalf("Migrate: postgres connect error: %s", err)
	}

	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		migrateLogger.Fatalf("Migrate: error apply migration: %v", err)
	}
	migrateLogger.Println("Migrate: all migrations applied")
}
