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

func GetWorkspaces(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := uuid.Parse(c.GetString("userID"))
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError))
			return
		}

		workspaceRepo := repositories.NewWorkspaceRepository(db.DB)
		workspaceService := services.NewWorkspaceService(workspaceRepo)

		workspaces, err := workspaceService.GetWorkspaces(userID)
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Failed to get workspaces", http.StatusInternalServerError))
			return
		}

		HandleResponse(c, http.StatusOK, "Workspaces retrieved", workspaces)
	}
}
