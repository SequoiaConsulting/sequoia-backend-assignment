package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardUserRepository implements boardUser CRUD interface
type BoardUserRepository interface {
	GetByID(string) (*model.BoardUser, error)
	Create(*model.BoardUser) error
	Update(*model.BoardUser) error
	Delete(*model.BoardUser) error
}
