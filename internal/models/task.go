package model

import (
	"time"
)

type Task struct {
	Base
	Title       string
	Description string
	DueDate     time.Time
	Status      string
	Priority    string
	CompletedAt time.Time
	CategoryID  uint
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	WorkspaceID uint
	Workspace   Workspace `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint
	User        User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Assignments []TaskAssignment `gorm:"foreignKey:TaskID"`
	Reminders   []Reminder       `gorm:"foreignKey:TaskID"`
}
