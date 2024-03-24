package repositories

import (
	models "github.com/StefanWellhoner/task-manager-api/internal/models"
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
