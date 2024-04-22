package model

import "github.com/google/uuid"

type Role struct {
	Base
	Name        string       `gorm:"not null" json:"name"`
	Description string       `json:"description"`
	WorkspaceID uuid.UUID    `gorm:"type:uuid;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
