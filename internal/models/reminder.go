package model

import (
	"time"
)

type Reminder struct {
	Base
	ReminderDate time.Time
	TaskID       uint
	Task         Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
