package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// TaskMySQLRepository stores database connection
type TaskMySQLRepository struct {
	db *gorm.DB
}

// NewTaskMySQLRepository accepts TaskMySQLRepository and returns TaskMySQLRepository
func NewTaskMySQLRepository(db *gorm.DB) *TaskMySQLRepository {
	return &TaskMySQLRepository{db}
}

// GetByID does a database query to get a task by ID
func (repo *TaskMySQLRepository) GetByID(id string) (*model.Task, error) {
	task := &model.Task{}
	if err := repo.db.Where("id = ?", id).First(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// Create does a database query to create a task
func (repo *TaskMySQLRepository) Create(task *model.Task) error {
	if err := repo.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// Update does a database query to update a task by ID
func (repo *TaskMySQLRepository) Update(task *model.Task) error {

	if err := repo.db.Update(task).Error; err != nil {
		return err
	}
	return nil
}

// Delete does a database query to delete a task by ID
func (repo *TaskMySQLRepository) Delete(task *model.Task) error {
	if err := repo.db.Delete(task).Error; err != nil {
		return err
	}
	return nil
}
