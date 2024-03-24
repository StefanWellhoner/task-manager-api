package handlers

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

func GetUser(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRepo := repositories.NewTokenRepository(db.DB)
		userRepo := repositories.NewUserRepository(db.DB)
		tokenService := services.NewTokenService(tokenRepo)
		userService := services.NewUserService(userRepo, tokenService)

		user, err := userService.GetUserByID(c.Param("id"))
		if err != nil {
			HandleError(c, err)
			return
		}

		c.JSON(200, user.PublicProfile())
	}
}

func GetUsers(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func UpdateUser(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func DeleteUser(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}
