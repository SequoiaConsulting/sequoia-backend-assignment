package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardMySQLRepository stores database connection 
type BoardMySQLRepository struct {
	conn *gorm.DB
}

// NewBoardMySQLRepository accepts BoardMySQLRepository and returns BoardMySQLRepository
func NewBoardMySQLRepository(conn *gorm.DB) *BoardMySQLRepository {
	return &BoardMySQLRepository{conn}
}

// GetByID does a database query to get a board by ID
func (ur *BoardMySQLRepository) GetByID(id string) (*model.Board, error) {
	db := ur.conn
	board := &model.Board{}
	if err := db.Where("id = ?", id).First(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

// Create does a database query to create a board
func (ur *BoardMySQLRepository) Create(fields *model.Board, c *gin.Context) (*model.Board, error) {
	db := ur.conn
	board := &model.Board{}
	// Validate input
	if err := c.ShouldBindJSON(board); err != nil {
		return nil, err
	}
	db.Create(fields)
	return board, nil
}

// Update does a database query to update a board by ID
func (ur *BoardMySQLRepository) Update(id string, c *gin.Context) (*model.Board, error) {
	db := ur.conn
	board := &model.Board{}
	if err := db.Where("id = ?", c.Param("id")).Update(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

// Delete does a database query to delete a board by ID
func (ur *BoardMySQLRepository) Delete(id string, c *gin.Context) (*model.Board, error) {
	db := ur.conn
	board := &model.Board{}
	if err := db.Where("id = ?", c.Param("id")).First(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}
