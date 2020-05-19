package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardRepository implements board CRUD interface
type BoardRepository interface {
	GetByID(string) (*model.Board, error)
	Create(*model.Board) error
	Update(*model.Board) error
	Delete(*model.Board) error
}
