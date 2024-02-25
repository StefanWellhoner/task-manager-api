package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/StefanWellhoner/task-manager-api/internal/config"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(config *config.Configuration) (*GormDatabase, error) {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Database,
		config.Database.SSLMode)

	var db *gorm.DB
	var err error

	for i := 0; i < config.Database.MaxRetries; i++ {

		db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retry in %d seconds...", config.Database.RetryInterval)
		time.Sleep(time.Duration(config.Database.RetryInterval) * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object from GORM db: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(0)

	if err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return nil, fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	if err := db.AutoMigrate(
		new(model.User),
		new(model.Task),
		new(model.TaskAssignment),
		new(model.Category),
		new(model.PasswordResetToken),
		new(model.Permission),
		new(model.RefreshToken),
		new(model.Reminder),
		new(model.Role),
		new(model.Workspace),
		new(model.WorkspaceRole),
	); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	log.Printf("Connected to database %s on %s:%d", config.Database.Database, config.Database.Host, config.Database.Port)

	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) Close(ctx context.Context) error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get generic database object from GORM db: %w", err)
	}

	done := make(chan error)
	go func() {
		done <- sqlDB.Close()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

func (d *GormDatabase) Ping(ctx context.Context) error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get generic database object from GORM db: %w", err)
	}

	done := make(chan error)
	go func() {
		done <- sqlDB.Ping()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
