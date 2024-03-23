package model

import (
	"errors"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
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

type UserPublicProfile struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"profileImage"`
}

type UserPrivateProfile struct {
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	ProfileImage string    `json:"profileImage"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (u *User) VerifyPassword(password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

func (u *User) PublicProfile() *UserPublicProfile {
	return &UserPublicProfile{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		ProfileImage: u.ProfileImage,
	}
}

func (u *User) PrivateProfile() *UserPrivateProfile {
	return &UserPrivateProfile{
		Email:        u.Email,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		ProfileImage: u.ProfileImage,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidatePassword(password string) error {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	minLength := 8

	if len(password) < minLength {
		return errors.New("password must be at least 8 characters long")
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !(hasUpper && hasLower && hasNumber && hasSpecial) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

// TableName sets the table name for the user model.
func (User) TableName() string {
	return "users"
}
