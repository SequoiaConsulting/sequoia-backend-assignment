package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// StatusMySQLRepository stores database connection
type StatusMySQLRepository struct {
	db *gorm.DB
}

// NewStatusMySQLRepository accepts StatusMySQLRepository and returns StatusMySQLRepository
func NewStatusMySQLRepository(db *gorm.DB) *StatusMySQLRepository {
	return &StatusMySQLRepository{db}
}

// GetByID does a database query to get a status by ID
func (repo *StatusMySQLRepository) GetByID(id string) (*model.Status, error) {
	status := &model.Status{}
	if err := repo.db.Where("id = ?", id).First(status).Error; err != nil {
		return nil, err
	}
	return status, nil
}

// Create does a database query to create a status
func (repo *StatusMySQLRepository) Create(status *model.Status) error {

	if err := repo.db.Create(status).Error; err != nil {
		return err
	}
	return nil
}

// Update does a database query to update a status by ID
func (repo *StatusMySQLRepository) Update(status *model.Status) error {
	if err := repo.db.Where("id = ?", status.ID).Update(status).Error; err != nil {
		return err
	}
	return nil
}

// Delete does a database query to delete a status by ID
func (repo *StatusMySQLRepository) Delete(status *model.Status) error {
	if err := repo.db.Delete(status).Error; err != nil {
		return err
	}
	return nil
}
