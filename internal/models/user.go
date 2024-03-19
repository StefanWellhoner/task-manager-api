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
	Email               string               `gorm:"uniqueIndex" json:"email"`
	PasswordHash        string               `json:"password"`
	PasswordLastChanged time.Time            `json:"passwordLastChanged"`
	Verified            bool                 `gorm:"default:false" json:"verified"`
	FirstName           string               `json:"firstName"`
	LastName            string               `json:"lastName"`
	ProfileImage        string               `json:"profileImage"`
	LastLogin           time.Time            `json:"lastLogin"`
	Tasks               []Task               `gorm:"foreignKey:UserID" json:"tasks"`
	OwnedWorkspaces     []Workspace          `gorm:"foreignKey:OwnerUserID" json:"ownedWorkspaces"`
	WorkspaceRoles      []WorkspaceRole      `gorm:"many2many:workspace_roles;" json:"workspaceRoles"`
	RefreshTokens       []RefreshToken       `gorm:"foreignKey:UserID" json:"refreshTokens"`
	PasswordResetTokens []PasswordResetToken `gorm:"foreignKey:UserID" json:"passwordResetTokens"`
}

type SanitizeUser struct {
	Base
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	ProfileImage string    `json:"profileImage"`
	LastLogin    time.Time `json:"lastLogin"`
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

func (u *User) SanitizeUser() *SanitizeUser {
	return &SanitizeUser{
		Base:         u.Base,
		Email:        u.Email,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		ProfileImage: u.ProfileImage,
		LastLogin:    u.LastLogin,
	}
}

// TableName sets the table name for the user model.
func (User) TableName() string {
	return "users"
}
