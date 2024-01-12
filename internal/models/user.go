package model

import (
	"time"

	"gorm.io/gorm"
)

// User Model
//
// User model is used to store user information.
//
// swagger:model User
type User struct {
	// The user ID
	//
	// required: true
	// example: 1
	gorm.Model
	// The user username
	//
	// required: true
	// example: johndoe
	Username string `gorm:"type:varchar(180);unique_index"`
	// The user password
	//
	// required: true
	// example: $2a
	Password string `gorm:"type:varchar(180)"`
	// The user email
	//
	// required: true
	// example: johndoe@example.com
	Email string `gorm:"type:varchar(180);unique_index"`
	// The user first name
	//
	// required: true
	// example: John
	FirstName string `gorm:"type:varchar(180)"`
	// The user last name
	//
	// required: true
	// example: Doe
	LastName string `gorm:"type:varchar(180)"`
	// The user is active
	//
	// required: true
	// example: true
	IsActive bool `gorm:"type:boolean;default:true"`
	// The user's last login date/time
	//
	// required: true
	// example: 2020-01-01T00:00:00Z
	LastLogin time.Time
	// A slice of tasks associated with this user
	//
	// required: true
	// example: ["1","2","3"]
	Tasks []Task `gorm:"foreignKey:UserID"`
}

// TableName sets the table name for the user model.
func (User) TableName() string {
	return "users"
}
