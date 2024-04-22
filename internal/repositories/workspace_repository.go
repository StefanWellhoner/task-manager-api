package repositories

import (
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkspaceRepository is the repository for the workspace model.
type WorkspaceRepository struct {
	db *gorm.DB
}

// NewWorkspaceRepository creates a new workspace repository.
func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{db: db}
}

// Create creates a new workspace in the database.
func (r *WorkspaceRepository) Create(workspace *models.Workspace) error {
	tx := r.db.Begin()

	if err := tx.Create(workspace).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *WorkspaceRepository) GetWorkspaces(userID uuid.UUID) ([]models.Workspace, error) {
	var workspaces []models.Workspace
	workspaceIDs := []uuid.UUID{}

	r.db.Model(&models.Workspace{}).
		Select("workspace.id").
		Joins("JOIN workspace_members ON workspace_members.workspace_id = workspace.id").
		Where("workspace_members.user_id = ?", userID).
		Find(&workspaceIDs)

	err := r.db.Model(&models.Workspace{}).
		Where("owner_id = ?", userID).
		Or("id IN ?", workspaceIDs).
		Find(&workspaces).Error

	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (r *WorkspaceRepository) GetWorkspace(worspaceID uuid.UUID) (*models.Workspace, error) {
	var workspace models.Workspace

	if err := r.db.Where("id = ?", worspaceID).First(&workspace).Error; err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (r *WorkspaceRepository) DeleteWorkspace(worspaceID uuid.UUID) error {
	tx := r.db.Begin()

	if err := tx.Where("id = ?", worspaceID).Delete(&models.Workspace{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
