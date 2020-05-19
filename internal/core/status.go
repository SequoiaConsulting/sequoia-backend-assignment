package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

// StatusCore struct to implement other menthods of status
type StatusCore struct {
	repo repository.StatusRepository
}

// NewStatusCore accepts repository.StatusRepository and returns StatusCore
func NewStatusCore(repo repository.StatusRepository) *StatusCore {
	return &StatusCore{repo}
}

// GetByID is the core domain layer method to get status by the ID
func (core *StatusCore) GetByID(id string) (*model.Status, error) {
	status, err := core.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return status, err
}

// GetByID is the core domain layer method to get status by the name
func (core *StatusCore) GetByName(id string) (*model.Status, error) {
	status, err := core.repo.GetByName(id)
	if err != nil {
		return nil, err
	}
	return status, err
}

// Create is the core domain layer method to create a status
func (core *StatusCore) Create(status *model.Status) (*model.Status, error) {
	if err := core.repo.Create(status); err != nil {
		return nil, err
	}
	return status, nil
}

// Update is the core domain layer method to update a status
func (core *StatusCore) Update(id string, status *model.Status) (*model.Status, error) {
	// check if status.ID  == null -> throw custom error
	if err := core.repo.Update(id, status); err != nil {
		return nil, err
	}
	return status, nil
}

// Delete is the core domain layer method to delete a status
func (core *StatusCore) Delete(status *model.Status) (*model.Status, error) {
	if err := core.repo.Delete(status); err != nil {
		return nil, err
	}
	return status, nil
}
