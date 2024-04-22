package dto

import (
	"time"

	models "github.com/StefanWellhoner/task-manager-api/internal/models"
)

type PublicUserDTO struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	ProfilePicture string `json:"profilePicture"`
}

type PrivateUserDTO struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Verified       bool      `json:"verified"`
	ProfilePicture string    `json:"profilePicture"`
	LastLogin      time.Time `json:"lastLogin"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func ToUserPrivateDTO(user *models.User) *PrivateUserDTO {
	return &PrivateUserDTO{
		ID:             user.ID.String(),
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Verified:       user.Verified,
		ProfilePicture: user.ProfileImage,
		LastLogin:      user.LastLogin,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func ToUserPublicDTO(user *models.User) *PublicUserDTO {
	return &PublicUserDTO{
		ID:             user.ID.String(),
		ProfilePicture: user.ProfileImage,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
	}
}
