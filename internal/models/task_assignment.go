package model

import (
	"time"
)

type TaskAssignment struct {
	Base
	AssignedAt time.Time
	UserID     uint
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TaskID     uint
	Task       Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
