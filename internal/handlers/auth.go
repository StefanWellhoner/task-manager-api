package handlers

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/dto"
	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repositories"
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

func setAuthCookie(c *gin.Context, tokenDetails *services.TokenDetails) {
	c.SetCookie("access_token", tokenDetails.AccessToken, 60*15, "/", "", true, true)
	c.SetCookie("refresh_token", tokenDetails.RefreshToken, 60*60*24*7, "/", "", true, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

func revokeAuthCookie(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

func Login(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LoginRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid payload", http.StatusBadRequest))
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

		setAuthCookie(c, tokenDetails)

		HandleResponse(c, http.StatusOK, "Login successful", nil)
	}
}

func Refresh(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Refresh token not found", http.StatusBadRequest))
			return
		}

		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)

		tokenDetails, err := tokenService.RefreshToken(refreshToken)
		if err != nil {
			HandleError(c, err)
			return
		}

		setAuthCookie(c, tokenDetails)

		HandleResponse(c, http.StatusOK, "Token refreshed", nil)
	}
}

func Logout(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")

		if err == nil {
			tokenRepo := repositories.NewTokenRepository(db.DB)
			tokenService := services.NewTokenService(tokenRepo)
			_ = tokenService.DeleteRefreshToken(refreshToken)
		}

		revokeAuthCookie(c)

		HandleResponse(c, http.StatusOK, "Logout successful", nil)
	}
}

func ChangePassword(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload ChangePasswordRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid payload", http.StatusBadRequest))
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
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func ResetPasswordConfirm(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func VerifyEmail(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func GetUserFromToken(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")

		userRepo := repositories.NewUserRepository(db.DB)
		tokenRepo := repositories.NewTokenRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		user, err := userService.GetUserByID(userID.(string))
		if err != nil {
			HandleError(c, err)
			return
		}
		HandleResponse(c, http.StatusOK, "User found", dto.ToUserPrivateDTO(user))
	}
}

func Register(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RegisterRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid payload", http.StatusBadRequest))
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

		HandleResponse(c, http.StatusCreated, "User registered successfully", dto.ToUserPrivateDTO(&user))
	}
}
