package database

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"raz/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrMissingMigrationsPath = errors.New("MIGRATIONS_PATH env missing")
	ErrMissingDatabaseURL    = errors.New("DATABASE_URL env missing")
)

func loadConfigFromURL() (*pgxpool.Config, error) {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, fmt.Errorf("Must set DATABASE_URL env var")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return config, nil
}

func loadConfig() (*pgxpool.Config, error) {
	cfg, err := config.NewDatabase()
	if err != nil {
		return loadConfigFromURL()
	}

	return pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	))
}

func dbURL() (string, error) {
	cfg, err := config.NewDatabase()
	if err != nil {
		dbURL, ok := os.LookupEnv("DATABASE_URL")
		if !ok {
			return "", fmt.Errorf("Must set DATABASE_URL env var")
		}

		return dbURL, nil
	}

	return cfg.URL(), nil
}

func Connect(ctx context.Context, logger *slog.Logger, migrations fs.FS) (*pgxpool.Pool, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	logger.Debug("Running migrations")

	url, err := dbURL()
	if err != nil {
		return nil, err
	}

	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, fmt.Errorf("create source: %w", err)
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", source, url)
	if err != nil {
		return nil, fmt.Errorf("migrate new: %s", err)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return conn, nil
}
