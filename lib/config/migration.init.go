package config

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"                     // source file
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // driver PostgreSQL
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type Migration struct {
	cfg     Postgres
	migrate *migrate.Migrate
}

func NewMigration(cfg Postgres) *Migration {
	return &Migration{cfg: cfg}
}

func (m *Migration) MigrationInit() error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=Asia/Jakarta",
		m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.Database, m.cfg.Ssl)

	// Migrate path in folder file/migrations
	migrationsPath := "file://file/migrations"

	// Inisialisasi migrasi
	conn, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		return err
	}
	m.migrate = conn
	return nil
}

func (m *Migration) Run() {
	defer func(migrate *migrate.Migrate) {
		err, _ := migrate.Close()
		if err != nil {

		}
	}(m.migrate)

	if err := m.migrate.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func (m *Migration) Rollback() {
	defer func(migrate *migrate.Migrate) {
		err, _ := migrate.Close()
		if err != nil {

		}
	}(m.migrate)

	if err := m.migrate.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to rollback migrations: %v", err)
	}
}
