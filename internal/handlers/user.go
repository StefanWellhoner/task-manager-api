package handlers

import (
	"net/http"

	dto "github.com/StefanWellhoner/task-manager-api/internal/dto"
	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repositories"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get details of a user
// @Description Get details of a user by user ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id     path      string     true  "User ID"
// @Success 200 {object} dto.PublicUserDTO
// @Router /users/{id} [get]
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

		c.JSON(200, dto.ToUserPublicDTO(user))
	}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.PublicUserDTO
// @Router /users [get]
func GetUsers(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id     path      string     true  "User ID"
// @Param   user   body      dto.PublicUserDTO  true  "User object"
// @Success 200 {object} dto.PublicUserDTO
// @Router /users/{id} [put]
func UpdateUser(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id     path      string     true  "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUser(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}
