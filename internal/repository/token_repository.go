package repositories

import (
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(token *model.RefreshToken) error
	FindByToken(token string) (*model.RefreshToken, error)
	Delete(token string) error
}

type GormTokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &GormTokenRepository{DB: db}
}

func (r *GormTokenRepository) Create(token *model.RefreshToken) error {
	result := r.DB.Create(token)
	return result.Error
}

func (r *GormTokenRepository) FindByToken(token string) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	result := r.DB.Where("token = ?", token).First(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &refreshToken, nil
}

func (r *GormTokenRepository) FindByUserID(userID string) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	result := r.DB.Where("user_id = ?", userID).First(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &refreshToken, nil
}

func (r *GormTokenRepository) Delete(token string) error {
	result := r.DB.Where("token = ?", token).Delete(&model.RefreshToken{})
	return result.Error
}
