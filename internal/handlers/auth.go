package handlers

import (
	"net/http"

	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
}

func Refresh(c *gin.Context) {
}

func Logout(c *gin.Context) {
}

func Register(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RegisterRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			errors := make(map[string]string)
			for _, fieldError := range err.(validator.ValidationErrors) {
				errors[fieldError.Field()] = fieldError.ActualTag()
			}
			HandleError(c, http.StatusBadRequest, "Invalid regisration information", errors)
			return
		}

		if isRegistered(payload.Email, db) {
			HandleError(c, http.StatusBadRequest, "User already registered", nil)
			return
		}

		user := model.User{Email: payload.Email, PasswordHash: payload.Password, FirstName: payload.Firstname, LastName: payload.Lastname}

		result := db.DB.Create(&user)
		if result.Error != nil {
			HandleError(c, http.StatusInternalServerError, "Failed to register user", nil)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email, "firstname": user.FirstName, "lastname": user.LastName})
	}
}

func isRegistered(email string, db *services.GormDatabase) bool {
	var user model.User
	db.DB.Where("email = ?", email).First(&user)
	return user.ID != [16]byte{}
}
