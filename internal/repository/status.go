package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// StatusRepository implements status CRUD interface
type StatusRepository interface {
	GetByID(string) (*model.Status, error)
	Create(*model.Status) error
	Update(*model.Status) error
	Delete(*model.Status) error
}
