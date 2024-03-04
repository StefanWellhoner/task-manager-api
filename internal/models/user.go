package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User Model
//
// User model is used to store user information.
//
// swagger:model User
type User struct {
	Base
	Email               string `gorm:"uniqueIndex"`
	PasswordHash        string
	PasswordLastChanged time.Time
	Verified            bool `gorm:"default:false"`
	FirstName           string
	LastName            string
	ProfileImage        string
	LastLogin           time.Time
	Tasks               []Task               `gorm:"foreignKey:UserID"`
	OwnedWorkspaces     []Workspace          `gorm:"foreignKey:OwnerUserID"`
	WorkspaceRoles      []WorkspaceRole      `gorm:"many2many:workspace_roles;"`
	RefreshTokens       []RefreshToken       `gorm:"foreignKey:UserID"`
	PasswordResetTokens []PasswordResetToken `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	if len(u.PasswordHash) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = string(hashedPassword)
	}
	return
}

func (u *User) VerifyPassword(password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

// TableName sets the table name for the user model.
func (User) TableName() string {
	return "users"
}
