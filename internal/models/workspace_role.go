package model

import "github.com/google/uuid"

type WorkspaceRole struct {
	Base
	WorkspaceID uuid.UUID
	UserID      uuid.UUID
	RoleID      uuid.UUID
	Workspace   Workspace `gorm:"forgeinKey:WorkspaceID;"`
	User        User      `gorm:"forgeinKey:UserID;"`
	Role        Role      `gorm:"forgeinKey:RoleID;"`
}
