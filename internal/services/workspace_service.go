package services

import (
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
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
