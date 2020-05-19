package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

// TaskCore struct to implement other menthods of task
type TaskCore struct {
	repo repository.TaskRepository
}

// NewTaskCore accepts repository.TaskRepository and returns TaskCore
func NewTaskCore(repo repository.TaskRepository) *TaskCore {
	return &TaskCore{repo}
}

// GetByID is the core domain layer method to get task by the ID
func (core *TaskCore) GetByID(id string) (*model.Task, error) {
	task, err := core.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return task, err
}

// Create is the core domain layer method to create a task
func (core *TaskCore) Create(task *model.Task) (*model.Task, error) {
	if err := core.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

// Update is the core domain layer method to update a task
func (core *TaskCore) Update(task *model.Task) (*model.Task, error) {
	if err := core.repo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

// Delete is the core domain layer method to delete a task
func (core *TaskCore) Delete(task *model.Task) (*model.Task, error) {
	if err := core.repo.Delete(task); err != nil {
		return nil, err
	}
	return task, nil
}
