package repositories

import (
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Update(user *model.User) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user *model.User) error {
	result := r.DB.Create(user)
	return result.Error
}

func (r *GormUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) Update(user *model.User) error {
	result := r.DB.Save(user)
	return result.Error
}
