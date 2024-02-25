package model

type Workspace struct {
	Base
	Title       string
	Description string
	OwnerUserID uint
	Users       []WorkspaceRole `gorm:"foreignKey:WorkspaceID"`
	Tasks       []Task          `gorm:"foreignKey:WorkspaceID"`
}
