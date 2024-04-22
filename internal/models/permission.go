package model

type Permission struct {
	Base
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;"`
}
