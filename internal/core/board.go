package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
)

// BoardCore struct to implement other menthods of board
type BoardCore struct {
	repo repository.BoardRepository
}

// NewBoardCore accepts repository.BoardRepository and returns BoardCore
func NewBoardCore(repo repository.BoardRepository) *BoardCore {
	return &BoardCore{repo}
}

// GetByID is the core domain layer method to get board by the ID
func (core *BoardCore) GetByID(id string) (*model.Board, error) {
	board, err := core.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return board, err
}

// GetByID is the core domain layer method to get board by the name
func (core *BoardCore) GetByName(id string) (*model.Board, error) {
	board, err := core.repo.GetByName(id)
	if err != nil {
		return nil, err
	}
	return board, err
}

// Create is the core domain layer method to create a board
func (core *BoardCore) Create(board *model.Board) (*model.Board, error) {
	if err := core.repo.Create(board); err != nil {
		return nil, err
	}
	return board, nil
}

// Update is the core domain layer method to update a board
func (core *BoardCore) Update(id string, board *model.Board) (*model.Board, error) {
	// check if board.ID  == null -> throw custom error
	if err := core.repo.Update(id, board); err != nil {
		return nil, err
	}
	return board, nil
}

// Delete is the core domain layer method to delete a board
func (core *BoardCore) Delete(board *model.Board) (*model.Board, error) {
	if err := core.repo.Delete(board); err != nil {
		return nil, err
	}
	return board, nil
}


func (core *BoardCore) BoardUser(id string, board *model.Board, user *model.User) (*model.Board, error) {
	if err := core.repo.BoardUser(id, board, user); err != nil {
		return nil, err
	}
	return board, nil
}


func (core *BoardCore) BoardStatus(id string, board *model.Board, user *model.User) (*model.Board, error) {
	if err := core.repo.BoardStatus(id, board, user); err != nil {
		return nil, err
	}
	return board, nil
}