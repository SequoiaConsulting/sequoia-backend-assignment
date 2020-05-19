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
	return board, nil
}

// Create is the core domain layer method to create a board
func (core *BoardCore) Create(board *model.Board) (*model.Board, error) {
	// check for duplicates
	if err := core.repo.Create(board); err != nil {
		// throw custom error
		return nil, err
	}
	//
	return board, nil
}

// Update is the core domain layer method to update a board
func (core *BoardCore) Update(board *model.Board) (*model.Board, error) {
	// check if board.ID  == null -> throw custom error
	// Check if board exists (dbBoard) if dbBoard == nil
	if err := core.repo.Update(board); err != nil {
		return nil, err
		// better throw a custom error failedupdate
	}

	//return dbBoard
	return board, nil
}

// Delete is the core domain layer method to delete a board
func (core *BoardCore) Delete(id int) error {
	board := &model.Board{ID: id}
	if err := core.repo.Delete(board); err != nil {
		return err
	}
	return nil
}
