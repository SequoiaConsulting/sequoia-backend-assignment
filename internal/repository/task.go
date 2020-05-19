package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// TaskRepository implements task CRUD interface
type TaskRepository interface {
	GetByID(string) (*model.Task, error)
	Create(*model.Task) error
	Update(*model.Task) error
	Delete(*model.Task) error
}
