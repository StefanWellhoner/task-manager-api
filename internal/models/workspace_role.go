package model

import "github.com/google/uuid"

type WorkspaceRole struct {
	Base
	WorkspaceID uuid.UUID
	Workspace   Workspace `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      uuid.UUID
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID      uuid.UUID
	Role        Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
