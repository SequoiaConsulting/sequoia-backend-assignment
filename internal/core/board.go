package core

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

// BoardCore struct to implement other menthods of board
type BoardCore struct {
	ur repository.BoardRepository
}

// NewBoardCore accepts repository.BoardRepository and returns BoardCore
func NewBoardCore(ur repository.BoardRepository) *BoardCore {
	return &BoardCore{ur}
}

// GetByID is the core domain layer method to get board by the ID
func (c *BoardCore) GetByID(id string) (*model.Board, error) {
	board, err := c.ur.GetByID(id)
	if err != nil {
		return nil, err
	}
	return board, err
}

// GetByName is the core domain layer method to get board by the name
func (c *BoardCore) GetByName(name string) (*model.Board, error) {
	board, err := c.ur.GetByName(name)
	if err != nil {
		return nil, err
	}
	return board, err
}

// Create is the core domain layer method to create a board
func (c *BoardCore) Create(fields *model.Board, context *gin.Context) (*model.Board, error) {
	board, err := c.ur.Create(fields, context)
	if err != nil {
		return nil, err
	}
	return board, err
}

// Update is the core domain layer method to update a board
func (c *BoardCore) Update(id string, context *gin.Context) (*model.Board, error) {
	board, err := c.ur.Update(id, context)
	if err != nil {
		return nil, err
	}
	return board, err
}

// Delete is the core domain layer method to delete a board
func (c *BoardCore) Delete(id string, context *gin.Context) (*model.Board, error) {
	board, err := c.ur.Delete(id, context)
	if err != nil {
		return nil, err
	}
	return board, err
}
