package model

type WorkspaceRole struct {
	Base
	WorkspaceID uint
	Workspace   Workspace `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID      uint
	Role        Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
