package model

import "github.com/google/uuid"

type Workspace struct {
	Base
	Title       string `gorm:"not null"`
	Description string
	OwnerUserID uuid.UUID       `gorm:"not null"`
	Users       []WorkspaceRole `gorm:"foreignKey:WorkspaceID"`
	Tasks       []Task          `gorm:"foreignKey:WorkspaceID"`
}
