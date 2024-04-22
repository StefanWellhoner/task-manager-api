package model

import "github.com/google/uuid"

type Workspace struct {
	Base
	Name        string `gorm:"not null"`
	Description string
	OwnerID     uuid.UUID `gorm:"type:uuid;not null;"`
	Owner       User      `gorm:"foreignKey:OwnerID;"`
	Roles       []Role    `gorm:"foreignKey:WorkspaceID;"`
	Tasks       []Task    `gorm:"foreignKey:WorkspaceID;"`
	Users       []User    `gorm:"many2many:workspace_members;"`
}
