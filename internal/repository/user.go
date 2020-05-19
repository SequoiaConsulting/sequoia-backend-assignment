package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// UserRepository implements user CRUD interface
type UserRepository interface {
	GetByID(string) (*model.User, error)
	// GetByName(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}
