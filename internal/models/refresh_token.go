package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Base
	Token     string    `gorm:"index;unique;not null;" json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;" json:"userId"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
