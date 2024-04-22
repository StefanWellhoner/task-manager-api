package services

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkspaceService is the service for the workspace model.
type WorkspaceService struct {
	workspaceRepo *repositories.WorkspaceRepository
}

// NewWorkspaceService creates a new workspace service.
func NewWorkspaceService(workspaceRepo *repositories.WorkspaceRepository) *WorkspaceService {
	return &WorkspaceService{workspaceRepo: workspaceRepo}
}

// Create creates a new workspace in the database.
func (s *WorkspaceService) Create(workspace *models.Workspace) error {
	return s.workspaceRepo.Create(workspace)
}

func (s *WorkspaceService) GetWorkspaces(userID uuid.UUID) ([]models.Workspace, error) {
	return s.workspaceRepo.GetWorkspaces(userID)
}

func (s *WorkspaceService) GetWorkspace(worspaceID uuid.UUID) (*models.Workspace, error) {
	workspace, err := s.workspaceRepo.GetWorkspace(worspaceID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewServiceError(errors.NotFoundError, "Workspace not found", http.StatusNotFound)
		}
		return nil, errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError)
	}
	return workspace, nil
}

func (s *WorkspaceService) DeleteWorkspace(worspaceID uuid.UUID) error {
	return s.workspaceRepo.DeleteWorkspace(worspaceID)
}
