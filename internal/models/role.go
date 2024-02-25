package model

type Role struct {
	Base
	Name           string
	Description    string
	WorkspaceRoles []WorkspaceRole `gorm:"foreignKey:RoleID"`
	Permissions    []Permission    `gorm:"many2many:role_permissions;"`
}
