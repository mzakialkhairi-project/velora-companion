// Package bootstrap provides application startup functionality.
package bootstrap

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mzakiaklhairi/velora/internal/shared"
)

// Database holds the database connection
type Database struct {
	DB *gorm.DB
}

// InitDatabase initializes the database connection
func InitDatabase(cfg *shared.Config) (*Database, error) {
	shared.Info("Connecting to PostgreSQL...")

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), gormConfig)
	if err != nil {
		shared.Error("Failed to connect to database", "error", err.Error())
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		shared.Error("Failed to get database instance", "error", err.Error())
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		shared.Error("Failed to ping database", "error", err.Error())
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	shared.Info("PostgreSQL connected successfully",
		"host", cfg.DBHost,
		"port", cfg.DBPort,
		"database", cfg.DBName,
	)

	return &Database{DB: db}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Ping checks if database is reachable
func (d *Database) Ping(ctx context.Context) error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
