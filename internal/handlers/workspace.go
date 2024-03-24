package handlers

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkspaceResponse struct {
	ID          uuid.UUID                `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Owner       models.UserPublicProfile `json:"owner"`
}

type CreateWorkspaceRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func CreateWorkspace(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload CreateWorkspaceRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid payload", http.StatusBadRequest))
			return
		}

		userID, err := uuid.Parse(c.GetString("userID"))
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError))
			return
		}

		workspaceRepo := repositories.NewWorkspaceRepository(db.DB)
		workspaceService := services.NewWorkspaceService(workspaceRepo)

		workspace := &models.Workspace{
			Title:       payload.Title,
			Description: payload.Description,
			OwnerUserID: userID,
		}

		if err := workspaceService.Create(workspace); err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Failed to create workspace", http.StatusInternalServerError))
			return
		}

		HandleResponse(c, http.StatusCreated, "Workspace created", workspace)
	}
}
