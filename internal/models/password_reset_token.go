package model

import (
	"time"
)

type PasswordResetToken struct {
	Base
	Token     string
	ExpiresAt time.Time
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
