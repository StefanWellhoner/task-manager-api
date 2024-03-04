package handlers

import (
	"net/http"
	"time"

	"github.com/StefanWellhoner/task-manager-api/internal/config"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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

var (
	SigningKey         = []byte(config.Get().Secrets.Jwt)
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func Login(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LoginRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			errors := make(map[string]string)
			for _, fieldError := range err.(validator.ValidationErrors) {
				errors[fieldError.Field()] = fieldError.ActualTag()
			}
			HandleResponse(c, http.StatusBadRequest, "Invalid login information", errors)
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		userService := services.NewUserService(userRepo)

		tokenDetails, err := userService.Login(payload.Email, payload.Password)
		if err != nil {
			HandleResponse(c, http.StatusUnauthorized, "Invalid login information", nil)
			return
		}

		HandleResponse(c, http.StatusOK, "Login successful", gin.H{"access_token": tokenDetails.AccessToken, "refresh_token": tokenDetails.RefreshToken})
	}
}

func Refresh(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func Logout(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleResponse(c, http.StatusOK, "Logout successful", nil)
	}
}

func ChangePassword(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ResetPassword(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ResetPasswordConfirm(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func VerifyEmail(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetUserFromToken(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func Register(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RegisterRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			errors := make(map[string]string)
			for _, fieldError := range err.(validator.ValidationErrors) {
				errors[fieldError.Field()] = fieldError.ActualTag()
			}
			HandleResponse(c, http.StatusBadRequest, "Invalid registration information", errors)
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		userService := services.NewUserService(userRepo)

		user := model.User{Email: payload.Email, PasswordHash: payload.Password, FirstName: payload.Firstname, LastName: payload.Lastname}

		if err := userService.Register(&user); err != nil {
			if err.Error() == "user already exists" {
				HandleResponse(c, http.StatusConflict, "User is already registered", nil)
			} else {
				HandleResponse(c, http.StatusInternalServerError, "Failed to register user", nil)
			}
			return
		}

		HandleResponse(c, http.StatusCreated, "User registered successfully", gin.H{"id": user.ID, "email": user.Email, "firstname": user.FirstName, "lastname": user.LastName})
	}
}
