package database

import (
	"fmt"
	"hotel/internal/infrastructure/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

func MigrateUp(sourceURL string, config *config.Config) error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}
