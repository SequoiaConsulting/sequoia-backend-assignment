package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

// BoardUserCore struct to implement other menthods of boardUser
type BoardUserCore struct {
	repo repository.BoardUserRepository
}

// NewBoardUserCore accepts repository.BoardUserRepository and returns BoardUserCore
func NewBoardUserCore(repo repository.BoardUserRepository) *BoardUserCore {
	return &BoardUserCore{repo}
}

// GetByID is the core domain layer method to get boardUser by the ID
func (core *BoardUserCore) GetByID(id string) (*model.BoardUser, error) {
	boardUser, err := core.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return boardUser, err
}

// Create is the core domain layer method to create a boardUser
func (core *BoardUserCore) Create(boardUser *model.BoardUser) (*model.BoardUser, error) {
	if err := core.repo.Create(boardUser); err != nil {
		return nil, err
	}

	return boardUser, nil
}

// Update is the core domain layer method to update a boardUser
func (core *BoardUserCore) Update(boardUser *model.BoardUser) (*model.BoardUser, error) {
	if err := core.repo.Update(boardUser); err != nil {
		return nil, err
	}

	return boardUser, nil
}

// Delete is the core domain layer method to delete a boardUser
func (core *BoardUserCore) Delete(id int) error {
	board := &model.BoardUser{ID: id}
	if err := core.repo.Delete(board); err != nil {
		return err
	}
	return nil
}
