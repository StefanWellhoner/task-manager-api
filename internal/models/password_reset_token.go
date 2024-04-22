package model

import (
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	Base
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"index;unique;not null;"`
	ExpiresAt time.Time `json:"expiresAt"`
	User      User      `gorm:"foreignKey:UserID;"`
}
