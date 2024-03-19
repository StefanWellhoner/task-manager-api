package handlers

import (
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
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(200, user.SanitizeUser())
	}
}

func GetUsers(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}
