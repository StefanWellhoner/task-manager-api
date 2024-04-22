package dto

import (
	"time"

	models "github.com/StefanWellhoner/task-manager-api/internal/models"
)

type WorkspaceDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"ownerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToWorkspaceDTO(workspace *models.Workspace) *WorkspaceDTO {
	return &WorkspaceDTO{
		ID:          workspace.ID.String(),
		Name:        workspace.Name,
		Description: workspace.Description,
		OwnerID:     workspace.OwnerID.String(),
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
	}
}

func ToWorkspaceDTOs(workspaces []models.Workspace) []*WorkspaceDTO {
	var workspaceDTOs []*WorkspaceDTO
	for _, workspace := range workspaces {
		workspaceDTOs = append(workspaceDTOs, ToWorkspaceDTO(&workspace))
	}
	return workspaceDTOs
}
