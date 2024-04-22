package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Base
	Title       string
	Description string
	DueDate     time.Time
	Status      string
	Priority    string
	CompletedAt time.Time
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null;"`
	CreatorID   uuid.UUID  `gorm:"type:uuid;not null"`
	WorkspaceID uuid.UUID  `gorm:"type:uuid;not null"`
	Creator     User       `gorm:"foreignKey:CreatorID;"`
	Category    Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Workspace   Workspace  `gorm:"foreignKey:WorkspaceID"`
	Reminders   []Reminder `gorm:"foreignKey:TaskID"`
	Users       []User     `gorm:"many2many:user_tasks;"`
}
