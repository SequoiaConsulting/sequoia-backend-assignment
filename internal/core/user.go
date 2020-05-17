package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

// UserCore struct to implement other menthods of user
type UserCore struct {
	ur repository.UserRepository
}

// NewUserCore accepts repository.UserRepository and returns UserCore
func NewUserCore(ur repository.UserRepository) *UserCore {
	return &UserCore{ur}
}

// GetByID is the core domain layer method to get user by the ID
func (c *UserCore) GetByID(id string) (*model.User, error) {
	user, err := c.ur.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

// Create is the core domain layer method to create a user
func (c *UserCore) Create(fields *model.User, context *gin.Context) (*model.User, error) {
	user, err := c.ur.Create(fields, context)
	if err != nil {
		return nil, err
	}
	return user, err
}

// Update is the core domain layer method to update a user
func (c *UserCore) Update(id string, context *gin.Context) (*model.User, error) {
	user, err := c.ur.Update(id, context)
	if err != nil {
		return nil, err
	}
	return user, err
}

// Delete is the core domain layer method to delete a user
func (c *UserCore) Delete(id string, context *gin.Context) (*model.User, error) {
	user, err := c.ur.Delete(id, context)
	if err != nil {
		return nil, err
	}
	return user, err
}
