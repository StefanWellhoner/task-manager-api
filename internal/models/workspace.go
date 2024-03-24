package model

import "github.com/google/uuid"

type Workspace struct {
	Base
	Title       string `gorm:"not null"`
	Description string

	OwnerUserID uuid.UUID       `gorm:"not null"`
	Owner       User            `gorm:"foreignKey:OwnerUserID"`
	Users       []WorkspaceRole `gorm:"foreignKey:WorkspaceID"`
	Tasks       []Task          `gorm:"foreignKey:WorkspaceID"`
}
