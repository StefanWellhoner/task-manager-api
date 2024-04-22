package handlers

import (
	"net/http"

	dto "github.com/StefanWellhoner/task-manager-api/internal/dto"
	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repositories"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateWorkspaceRequest struct {
	Name        string `json:"name" binding:"required"`
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
			Name:        payload.Name,
			Description: payload.Description,
			OwnerID:     userID,
		}

		if err := workspaceService.Create(workspace); err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Failed to create workspace", http.StatusInternalServerError))
			return
		}

		HandleResponse(c, http.StatusCreated, "Workspace created", dto.ToWorkspaceDTO(workspace))
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

		HandleResponse(c, http.StatusOK, "Workspaces retrieved", dto.ToWorkspaceDTOs(workspaces))
	}
}

func GetWorkspace(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		workspaceID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid workspace ID", http.StatusBadRequest))
			return
		}

		workspaceRepo := repositories.NewWorkspaceRepository(db.DB)
		workspaceService := services.NewWorkspaceService(workspaceRepo)

		workspace, err := workspaceService.GetWorkspace(workspaceID)
		if err != nil {
			HandleError(c, err)
			return
		}

		HandleResponse(c, http.StatusOK, "Workspace retrieved", dto.ToWorkspaceDTO(workspace))
	}
}

func UpdateWorkspace(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleError(c, errors.NewServiceError(errors.NotImplemented, "Not implemented", http.StatusNotImplemented))
	}
}

func DeleteWorkspace(db *services.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		workspaceID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			HandleError(c, errors.NewServiceError(errors.ValidationError, "Invalid workspace ID", http.StatusBadRequest))
			return
		}

		workspaceRepo := repositories.NewWorkspaceRepository(db.DB)
		workspaceService := services.NewWorkspaceService(workspaceRepo)

		if err := workspaceService.DeleteWorkspace(workspaceID); err != nil {
			HandleError(c, errors.NewServiceError(errors.InternalError, "Failed to delete workspace", http.StatusInternalServerError))
			return
		}

		HandleResponse(c, http.StatusNoContent, "Workspace deleted", nil)
	}
}
