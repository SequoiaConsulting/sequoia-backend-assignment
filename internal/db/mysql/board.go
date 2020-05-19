package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardMySQLRepository stores database connection
type BoardMySQLRepository struct {
	db *gorm.DB
}

// NewBoardMySQLRepository accepts BoardMySQLRepository and returns BoardMySQLRepository
func NewBoardMySQLRepository(db *gorm.DB) *BoardMySQLRepository {
	return &BoardMySQLRepository{db}
}

// GetByID does a database query to get a board by ID
func (repo *BoardMySQLRepository) GetByID(id string) (*model.Board, error) {
	board := &model.Board{}
	fmt.Printf("%s", id)
	if err := repo.db.Where("id = ?", id).First(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

// GetByName does a database query to get a board by name
func (repo *BoardMySQLRepository) GetByName(id string) (*model.Board, error) {
	board := &model.Board{}
	if err := repo.db.Where("id = ?", id).First(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

// Create does a database query to create a board
func (repo *BoardMySQLRepository) Create(board *model.Board) error {
	if err := repo.db.Create(board).Error; err != nil {
		return err
	}
	return nil
}

// Update does a database query to update a board by ID
func (repo *BoardMySQLRepository) Update(id string, board *model.Board) error {
	if err := repo.db.Where("id = ?", id).Update(board).Error; err != nil {
		fmt.Printf("%v", board)
		return err
	}
	return nil
}

// Delete does a database query to delete a board by ID
func (repo *BoardMySQLRepository) Delete(board *model.Board) error {
	if err := repo.db.Delete(board).Error; err != nil {
		return err
	}
	return nil
}

// BoardUser 
func (repo *BoardMySQLRepository) BoardUser(id string, board_id *model.Board, user_id *model.User) error {
	if err := repo.db.Model(&model.Task{}).Where("board_id = ? AND assigned_to = ?", board_id, user_id).Error; err != nil {
		fmt.Printf("%v", board_id)
		return err
	}
	return nil
}

// BoardStatus
func (repo *BoardMySQLRepository) BoardStatus(id string, board_id *model.Board, user_id *model.User) error {
	if err := repo.db.Where("board_id = ? AND assigned_to = ?", board_id, user_id).Error; err != nil {
		fmt.Printf("%v", board_id)
		return err
	}
	return nil
}