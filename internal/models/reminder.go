package model

import (
	"time"

	"github.com/google/uuid"
)

type Reminder struct {
	Base
	ReminderDate time.Time `gorm:"not null" json:"reminderDate"`
	TaskID       uuid.UUID `gorm:"type:uuid;not null" json:"taskID"`
	Task         Task      `gorm:"foreignKey:TaskID;"`
}
