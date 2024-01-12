package services

import (
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(connection string) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)

	if err := db.AutoMigrate(new(model.User), new(model.Task)); err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *GormDatabase) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
