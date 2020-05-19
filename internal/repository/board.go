package repository

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardRepository implements board CRUD interface
type BoardRepository interface {
	GetByID(string) (*model.Board, error)
	GetByName(string) (*model.Board, error)
	Create(*model.Board) error
	Update(string, *model.Board) error
	Delete(*model.Board) error
	BoardUser(string, *model.Board, *model.User) error
	BoardStatus(string, *model.Board, *model.User) error
}
