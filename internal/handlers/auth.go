package handlers

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
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

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func Login(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LoginRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid login information", http.StatusBadRequest))
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		tokenDetails, err := userService.Authenticate(payload.Email, payload.Password)
		if err != nil {
			HandleError(c, err)
			return
		}

		HandleResponse(c, http.StatusOK, "Login successful", gin.H{"accessToken": tokenDetails.AccessToken, "refreshToken": tokenDetails.RefreshToken})
	}
}

func Refresh(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RefreshRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid refresh token", http.StatusBadRequest))
			return
		}

		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)

		

		tokenDetails, err := tokenService.RefreshToken(payload.RefreshToken)
		if err != nil {
			HandleError(c, err)
			return
		}
	}
}

func Logout(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LogoutRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid refresh token", http.StatusBadRequest))
			return
		}

		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)

		if err := tokenService.DeleteRefreshToken(payload.RefreshToken); err != nil {
			HandleError(c, err)
			return
		}

		HandleResponse(c, http.StatusOK, "Logout successful", nil)
	}
}

func ChangePassword(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload ChangePasswordRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid password change information", http.StatusBadRequest))
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		userID := c.GetString("userID")
		if err := userService.ChangePassword(userID, payload.OldPassword, payload.NewPassword); err != nil {
			HandleError(c, err)
			return
		}

		HandleResponse(c, http.StatusOK, "Password changed successfully", nil)
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
		userID, exists := c.Get("userID")
		if !exists {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError))
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		user, err := userService.GetUserByID(userID.(string))
		if err != nil {
			HandleError(c, err)
			return
		}
		HandleResponse(c, http.StatusOK, "User found", user.PrivateProfile())
	}
}

func Register(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RegisterRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid registration information", http.StatusBadRequest))
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		user := model.User{Email: payload.Email, PasswordHash: payload.Password, FirstName: payload.Firstname, LastName: payload.Lastname}

		if err := userService.Create(&user); err != nil {
			HandleError(c, err)
			return
		}

		HandleResponse(c, http.StatusCreated, "User registered successfully", user.PrivateProfile())
	}
}
