package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

// UserCore struct to implement other menthods of user
type UserCore struct {
	repo repository.UserRepository
}

// NewUserCore accepts repository.UserRepository and returns UserCore
func NewUserCore(repo repository.UserRepository) *UserCore {
	return &UserCore{repo}
}

// GetByID is the core domain layer method to get user by the ID
func (core *UserCore) GetByID(id string) (*model.User, error) {
	user, err := core.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

// Create is the core domain layer method to create a user
func (core *UserCore) Create(user *model.User) (*model.User, error) {
	if err := core.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Update is the core domain layer method to update a user
func (core *UserCore) Update(user *model.User) (*model.User, error) {
	if err := core.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete is the core domain layer method to delete a user
func (core *UserCore) Delete(id int) error {
	user := &model.User{ID: id}
	if err := core.repo.Delete(user); err != nil {
		return err
	}
	return nil
}
