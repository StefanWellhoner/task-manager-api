package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Base
	Token     string `gorm:"index;unique;not null;"`
	ExpiresAt time.Time
	UserID    uuid.UUID `gorm:"type:uuid;not null;"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
